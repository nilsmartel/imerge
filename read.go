package main

import (
	"image"
	"log"
	"os"
)

func readImage(path string) image.Image {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = reader.Close() }()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

