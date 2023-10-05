package main

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// Open the input image file
	inputFile, err := os.Open("./img/Ran Mitake_Phony Cover.png")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Decode the input image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	// Create an output image file
	outputFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Compression quality for JPEG (0-100)
	jpegOptions := jpeg.Options{Quality: 33} // Adjust quality as needed

	// Encode and save the compressed image
	err = jpeg.Encode(outputFile, img, &jpegOptions)
	if err != nil {
		panic(err)
	}

}
