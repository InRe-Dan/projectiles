package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	update(g *Game)
	render(s *ebiten.Image)
	getCollisionInfo() CollisionInfo
}

type CollisionInfo struct {
}

type CollisionBox struct {
}

type Wall struct {
	pos int
	dir Vect
}

func (w *Wall) update(g *Game) {
	// Repel player if collided with
}

func (w *Wall) render(s *ebiten.Image) {
	// No rendering done. Walls unseen.
}
