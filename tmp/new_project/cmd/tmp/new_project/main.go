package main
import (
	"log"

	eb "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sprite *eb.Image
}

func newGame() *Game {
	sprite := eb.NewImage(30, 30)
	return &Game{
		sprite,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) ClearBackground(screen *eb.Image) {
	_, _, size := getMeasures(screen)
	background := make([]byte, size)
	for i := 0; i < len(background); i += 4 {
		r := byte(100)
		g := byte(149)
		b := byte(237)
		background[i+0] = r
		background[i+1] = g
		background[i+2] = b
	}
	screen.WritePixels(background)
}

func (g *Game) Draw(screen *eb.Image) {
	g.ClearBackground(screen)
	_, _, size := getMeasures(g.sprite)
	pixels := make([]byte, size)
	for i := 0; i < size; i += 4 {
		alpha := 0.5
		r := byte(128)
		g := byte(128)
		b := byte(128)
		// rA := alphaIt(r, 100, alpha)
		// gA := alphaIt(g, 149, alpha)
		// bA := alphaIt(b, 237, alpha)
		pixels[i+0] = r
		pixels[i+1] = g
		pixels[i+2] = b
		pixels[i+3] = 255
	}
	g.sprite.WritePixels(pixels)
	screen.DrawImage(g.sprite, &eb.DrawImageOptions{})
}

func (g *Game) Layout(outSideWidth, outSideHeight int) (screenWidth, screenHeight int) {
	screenWidth = 320
	screenHeight = 240
	return
}

func main() {
	game := newGame()
	eb.SetWindowSize(game.Layout(0, 0))
	eb.SetWindowTitle("helloworld")
	if err := eb.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
