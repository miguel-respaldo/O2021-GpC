package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {

	imgPath := "cat.jpg"

	fmt.Println("Image dimmensions")
	fmt.Println(imgPath)

	file, err := os.Open(imgPath)

	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Width:", image.Width, "Height:", image.Height)
}
