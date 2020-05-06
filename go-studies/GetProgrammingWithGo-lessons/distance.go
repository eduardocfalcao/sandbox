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

func newLocation(name string, lat, long coordinate) location {
	return location{name, lat.decimal(), long.decimal()}
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

func main() {
	spirit := newLocation("Columbia Memorial Station", coordinate{d: 14, m: 34, s: 6.2, h: 'S'}, coordinate{d: 175, m: 28, s: 21.5, h: 'E'})
	opportunity := newLocation("Challenger Memorial Station", coordinate{d: 1, m: 56, s: 46.3, h: 'S'}, coordinate{d: 354, m: 28, s: 24.2, h: 'E'})
	curiosity := newLocation("Bradbury Landing", coordinate{d: 4, m: 35, s: 22.2, h: 'S'}, coordinate{d: 137, m: 26, s: 30.1, h: 'E'})
	inSight := newLocation("Elysium Planitia", coordinate{d: 4, m: 30, s: 0.0, h: 'N'}, coordinate{d: 135, m: 54, s: 0, h: 'E'})

	mars := world{3389.5}

	fmt.Printf("Distance between Spirit and Opportunity %.2f KM.\n", mars.distante(spirit, opportunity))
	fmt.Printf("Distance between Spirit and Curiosity %.2f KM.\n", mars.distante(spirit, curiosity))
	fmt.Printf("Distance between Spirit and InSight %.2f KM.\n", mars.distante(spirit, inSight))

	fmt.Printf("Distance between Opportunity and Curiosity %.2f KM.\n", mars.distante(opportunity, curiosity))
	fmt.Printf("Distance between Opportunity and InSight %.2f KM.\n", mars.distante(opportunity, inSight))

	fmt.Printf("Distance between Curiosity and InSight %.2f KM.\n", mars.distante(curiosity, inSight))

}
