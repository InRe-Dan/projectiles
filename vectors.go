package main

import (
	"math"
)

type Vect struct {
	x float64
	y float64
}

func (v1 *Vect) add(v2 *Vect) {
	v1.x += v2.x
	v1.y += v2.y
}

func (v *Vect) mult(scalar float64) {
	v.x *= scalar
	v.y *= scalar
}

func (v *Vect) mag() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vect) unitVect() *Vect {
	magnitude := v.mag()
	return &Vect{v.x / magnitude, v.y / magnitude}
}

func (v *Vect) limitMag(limit float64) {
	if v.mag() > limit {
		new := v.unitVect()
		new.mult(limit)
		v.x = new.x
		v.y = new.y
	}
}
