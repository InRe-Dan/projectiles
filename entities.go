package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Entity interface {
	update(g *Game)
	render(s *ebiten.Image)
	getCollisionInfo() *CollisionInfo
	determineCollision(info CollisionInfo) (bool, *Vect)
	isAlive() bool
}

type CollisionType int

const (
	static CollisionType = iota
	destructible
	moveable
)

type CollisionInfo struct {
	pos           Vect
	prevPos       Vect
	previousBoxes []CollisionBox
	boxes         []CollisionBox
	collisionType CollisionType
}

type CollisionBox struct {
	topLeft  Vect
	topRight Vect
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

func (w *Wall) getCollisionInfo() CollisionInfo {
	return CollisionInfo{}
}

func (w *Wall) determineCollision() (bool, Vect) {
	return true, Vect{}
}
func (w *Wall) isAlive() bool {
	return true
}

type Player struct {
	prevPos Vect
	pos     Vect
	vel     Vect
	image   *ebiten.Image
}

func (p *Player) render(s *ebiten.Image) {
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
			dv = dv.mult(BOOSTACCELERATION)
		} else {
			dv = dv.mult(PLAYERACCELERATION)
		}
		// Apply
		p.vel = *p.vel.add(dv)

	}

	// Apply drag to the player
	// Copy the direction of velocity, reverse it and multiply it
	velMag := p.vel.mag()
	dragVect := &Vect{p.vel.x, p.vel.y}
	dragVect = dragVect.unitVect()
	if moveDir.mag() > 0.0 {
		dragVect = dragVect.mult(-DRAGCOEFF * velMag * velMag)
	} else {
		dragVect = dragVect.mult(-STANDSTILLDRAGCOEFF * velMag * velMag)
	}
	p.vel = *p.vel.add(dragVect)

	p.pos = *p.pos.add(&p.vel)
}

func (p *Player) getCollisionInfo() *CollisionInfo {
	prevHitBox := CollisionBox{p.prevPos, *p.prevPos.copy().add(&Vect{10, 10})}
	prevHitBoxes := make([]CollisionBox, 1)
	prevHitBoxes = append(prevHitBoxes, prevHitBox)

	currentHitBox := CollisionBox{p.pos, *p.pos.copy().add(&Vect{10, 10})}
	currentHitBoxes := make([]CollisionBox, 1)
	currentHitBoxes = append(currentHitBoxes, currentHitBox)

	collisionInfo := CollisionInfo{p.pos, p.prevPos, prevHitBoxes, currentHitBoxes, moveable}
	return &collisionInfo
}
