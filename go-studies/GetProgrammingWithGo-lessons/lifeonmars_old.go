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
	right command = 0
	left  command = 1
)

type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex
	cells  [][]cell
}

type cell struct {
	occupier *Occupier
}

type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

type Message struct {
	Pos       image.Point
	LifeSigns int
	Rover     string
}

type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
}

func newMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}

	for i := range grid.cells {
		grid.cells[i] = make([]cell, size.X)
	}
	return grid
}

func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
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
		pos:  p,
	}

	return cell.occupier
}

func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()

	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}

	o.grid.cell(o.pos).occupier = nil
	newCell.occupier = o
	o.pos = p
	return true
}

func (o *Occupier) Pos() image.Point {
	return o.pos
}

func startDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	// Try a random point; continue until we've found one that's
	// not currently occupied.
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint)
	}
	return NewRoverDriver(name, o, marsToEarth)
}

func NewRoverDriver(
	name string,
	occupier *Occupier,
	marsToEarth chan []Message,
) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
		name:     name,
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.Pos())
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
			log.Printf("%s new direction %v", r.name, direction)
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s moved to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)
			// Pick one of the other directions randomly.
			// Next time round, we'll try to move in the new
			// direction.
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
