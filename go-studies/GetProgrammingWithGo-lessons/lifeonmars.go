package main

import (
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
)

type MarsGrid struct {
	mu     sync.Mutex
	bounds image.Rectangle
	grid   [][]cell
}

func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

func NewMarsGrid(size image.Point) *MarsGrid {
	g := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		grid: make([][]cell, size.Y),
	}
	for i := range g.grid {
		g.grid[i] = make([]cell, size.X)
		for j := range g.grid[i] {
			g.grid[i][j] = NewCell()
		}
	}
	return g
}

type ground struct {
	lifeSign bool
}

type cell struct {
	occupier *Occupier
	ground   ground
}

func NewCell() cell {
	return cell{
		ground: ground{
			lifeSign: rand.Intn(1000) >= 700,
		},
	}
}

func (c *cell) lifeSign() bool {
	return c.ground.lifeSign
}

type Occupier struct {
	grid *MarsGrid
	cell *cell
	pos  image.Point
}

func (o *Occupier) Pos() image.Point {
	return o.pos
}

func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.grid[p.Y][p.X]
}

func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)

	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		cell: cell,
		pos:  p,
	}
	return cell.occupier
}

func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	cell := o.grid.cell(p)
	if cell == nil || cell.occupier != nil {
		return false
	}

	o.grid.cell(o.pos).occupier = nil
	cell.occupier = o
	o.pos = p
	return true
}

type RoverDriver struct {
	name     string
	occupier *Occupier
	marsGrid *MarsGrid
	radio    *Radio
	commandc chan command
}

func (r *RoverDriver) drive() {
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				if r.occupier.cell.lifeSign() {
					r.radio.send(Message{
						coordinates: newPos,
						roverName:   r.name,
					})
				}
				log.Printf("%s moved to position %+v", r.name, newPos)
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)

			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)

		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func NewRoverDriver(marsGrid *MarsGrid, name string, o *Occupier, radio *Radio) *RoverDriver {
	r := &RoverDriver{
		name:     name,
		occupier: o,
		marsGrid: marsGrid,
		radio:    radio,
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

type Message struct {
	coordinates image.Point
	roverName   string
}
type Radio struct {
	messages chan Message
}

func NewRadio() *Radio {
	return &Radio{
		messages: make(chan Message),
	}
}

type Satellite struct {
	messages []Message
	radio    *Radio
	earth    *Earth
}

func NewSatellite(r *Radio, e *Earth) *Satellite {
	s := &Satellite{
		radio:    r,
		earth:    e,
		messages: make([]Message, 0),
	}

	go s.sendMessagesToEarth()
	return s
}

type Earth struct {
	lifeSignOnMars chan []Message
}

func NewEarth() *Earth {
	return &Earth{
		lifeSignOnMars: make(chan []Message),
	}
}

func (s *Satellite) canSendToEarth(timeOfTheDay int) bool {
	return timeOfTheDay > 0 && timeOfTheDay < 5
}

func (s *Satellite) sendMessagesToEarth() {
	timeOfTheDay := 0
	next := time.After(time.Second)
	for {
		select {
		case m := <-s.radio.messages:
			s.messages = append(s.messages, m)
		case <-next:
			if s.canSendToEarth(timeOfTheDay) {
				s.earth.lifeSignOnMars <- s.messages
				s.messages = nil
			}
			timeOfTheDay++
			if timeOfTheDay == 24 {
				timeOfTheDay = 0
			}
			next = time.After(time.Second)
		}
	}
}

func (e *Earth) reportLifeSign() {
	for messages := range e.lifeSignOnMars {
		log.Println("Messages received from Mars rovers:")
		for _, m := range messages {
			log.Printf("Rover %s found live in postition %v", m.roverName, m.coordinates)
		}
	}
}

func (r *Radio) send(m Message) {
	r.messages <- m
}

func main() {
	mars := NewMarsGrid(image.Point{
		X: 200,
		Y: 200,
	})

	radio := NewRadio()
	earth := NewEarth()
	NewSatellite(radio, earth)

	for _, roverName := range []string{"Sojourner", "Opportunity", "Spirit", "Curiosity", "Perseverance"} {
		var o *Occupier

		for o == nil {
			startPoint := image.Point{X: rand.Intn(mars.Size().X), Y: rand.Intn(mars.Size().Y)}
			o = mars.Occupy(startPoint)
		}
		NewRoverDriver(mars, roverName, o, radio)
	}
	earth.reportLifeSign()
}
