package img

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"log"
	"os"
)

func Read(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = reader.Close() }()
	img, _, err := image.Decode(reader)

	return img, err
}

