package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	pos   *Vect
	vel   *Vect
	image *ebiten.Image
}

func (p *Player) render(s ebiten.Image) {
	matrix := ebiten.GeoM{}
	matrix.Translate(p.pos.x, p.pos.y)
	s.DrawImage(p.image, &ebiten.DrawImageOptions{GeoM: matrix})
}

func (p *Player) update(g *Game) {
	// Determine direction of movement
	moveDir := &Vect{0, 0}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		moveDir.x += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		moveDir.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		moveDir.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		moveDir.y -= 1
	}

	// Accelerate the player
	if moveDir.mag() > 0.0 {
		// Normalise
		dv := moveDir.unitVect()
		// Multiply by acceleration
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			dv.mult(BOOSTACCELERATION)
		} else {
			dv.mult(PLAYERACCELERATION)
		}
		// Apply
		p.vel.add(dv)

	}

	// Apply drag to the player
	// Copy the direction of velocity, reverse it and multiply it
	velMag := p.vel.mag()
	dragVect := &Vect{p.vel.x, p.vel.y}
	dragVect.unitVect()
	if moveDir.mag() > 0.0 {
		dragVect.mult(-DRAGCOEFF * velMag * velMag)
	} else {
		dragVect.mult(-STANDSTILLDRAGCOEFF * velMag * velMag)
	}
	p.vel.add(dragVect)

	p.pos.add(p.vel)
}
