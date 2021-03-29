package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth, screenHeight = 640, 480
)

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// 背景を描写
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player{spaceShip, screenWidth / 2.0, screenHeight / 2.0, 4}
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

func update(screen *ebiten.Image) error {
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "game!"); err != nil {
		log.Fatal(err)
	}
}
