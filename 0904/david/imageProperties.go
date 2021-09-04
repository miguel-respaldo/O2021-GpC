package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func getImageProperties(filePath string) {
	fileImage, err := os.Open(filePath)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	imageData, _, err := image.DecodeConfig(fileImage)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println("wxh: ", imageData.Width, "x", imageData.Height)
}

func main() {
	getImageProperties("0904.png")
}
