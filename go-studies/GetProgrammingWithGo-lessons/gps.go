package main

import (
	"fmt"
	"math"
)

type location struct {
	Name string  `json:name`
	Lat  float64 `json:latitude`
	Long float64 `json:longitude`
}

func newLocation(name string, lat, long float64) location {
	return location{name, lat, long}
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type world struct {
	raiuds float64
}

func (w world) distante(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))
	return w.raiuds * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type gps struct {
	current, destination location
	world                world
}

func (g gps) distante() float64 {
	return g.world.distante(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("Remain %.2f Km to reach the destination", g.distante())
}

type rover struct {
	gps
}

func main() {
	mars := world{3389.5}
	breadbury := newLocation("Bradbury Landing", -4.5895, 137.4417)
	destination := newLocation("Elysium Planitia", 4.5, 135.9)
	gps := gps{current: breadbury, destination: destination, world: mars}

	curiosity := rover{gps: gps}

	fmt.Println(curiosity.message())

}
