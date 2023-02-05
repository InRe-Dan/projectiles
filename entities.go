package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	update(g Game)
	render(s ebiten.Image)
}

type Wall struct {
	colour color.Color
	pos    Vect
	size   Vect
}

func (w *Wall) update() {
	// Repel player if collided with

}

func (w *Wall) render() {
	// No rendering done. Walls unseen.
}
