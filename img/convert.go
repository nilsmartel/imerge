package img

import (
	"image"
)

func ToRGBA(i image.Image) image.RGBA {
	conv := i.ColorModel().Convert
	newImage := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{
			i.Bounds().Dx(),
			i.Bounds().Dy(),
			},
		})

	for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
		for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
			newImage.Set(x, y, conv(i.At(x, y)))
		}
	}

	return *newImage
}
