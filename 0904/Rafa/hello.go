package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// Read image from file that already exists
	existingImageFile, err := os.Open("C:/JPG.jpg")
	if err != nil {
		// No errors lol
	}

	defer existingImageFile.Close()

	m, _, err := image.Decode(existingImageFile)
	if err != nil {
		// No errors lol
	}

	g := m.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	// Print results
	fmt.Println("La imagen tiene un tamaño de:", width, "x", height)
}
