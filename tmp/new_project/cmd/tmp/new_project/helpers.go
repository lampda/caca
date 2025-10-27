package main
import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

func getMeasures(img *eb.Image) (int, int, int) {
	imgRect := img.Bounds()
	width := imgRect.Max.X - imgRect.Min.X
	height := imgRect.Max.Y - imgRect.Min.Y
	size := width * height * 4
	return width, height, size
}

func alphaIt(srcColor, dstColor byte, alpha float64) byte {
	src := (float64(srcColor) * alpha)
	dst := (float64(dstColor) * (1 - alpha))
	final := src + dst
	b := byte(final)
	return b
}
