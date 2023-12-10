// Sharome Burton

package main

import (
	"fmt"

	"github.com/koulkoudakis/img_mod/Colors"
	"github.com/koulkoudakis/img_mod/Grayscale"
	"github.com/koulkoudakis/img_mod/Text"
	"github.com/koulkoudakis/img_mod/getPic"
)

func main() {
	// Prints Heading of the Program
	fmt.Println("Sharome Burton")
	fmt.Println("Program 3 - Image Ascii Art")

	fmt.Print("\n")

	// Call function to get picture from URL
	getPic.DownloadPicture()

	fmt.Print("\n")

	// Call function to process Pixel Colors
	Colors.PrintPixels()

	fmt.Print("\n")

	// Call function to grayscale image
	Grayscale.GrayScale()

	fmt.Print("\n")

	// Call function to print colored text to image
	Text.PrintText()
}
