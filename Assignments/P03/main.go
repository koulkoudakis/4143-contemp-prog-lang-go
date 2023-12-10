// Sharome Burton

package main

import (
	"fmt"
	"github.com/koulkoudakis/img_mod/Colors"
	"github.com/koulkoudakis/img_mod/getPic"
	"github.com/koulkoudakis/img_mod/Grayscale"
	"github.com/koulkoudakis/img_mod/Text"
)

func main() {
	fmt.Println("Sharome Burton")
	fmt.Println("Program 3 - Image Ascii Art")

	// get picture from URL
	getPic.main()

	// process pixel colors
	Colors.main()

	// grayscale image
	Grayscale.main()

	// print colored text to image
	Text.main()
}
