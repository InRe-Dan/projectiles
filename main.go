package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	DRAGCOEFF           = 0.001
	STANDSTILLDRAGCOEFF = 0.001
	PLAYERACCELERATION  = 0.2
	BOOSTACCELERATION   = 10
)

type Game struct {
	p        *Player
	entities []Entity
}

func (g *Game) Update() error {
	g.p.update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.p.render(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%d, %d)\n(%f, %f)\n%f",
		int(g.p.pos.x), int(g.p.pos.y), (g.p.vel.x), (g.p.vel.y), (g.p.vel.mag())))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Project")
	g := &Game{}
	p := &Player{}
	g.p = p
	g.p.image = ebiten.NewImage(10, 10)
	g.p.image.Fill(color.RGBA{50, 50, 50, 255})
	g.p.pos = Vect{100, 100}
	g.p.vel = Vect{0, 0}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
