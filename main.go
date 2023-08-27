package main

import (
	"fmt"
	"os"
)

const ecranInspirantUrl = "get.ecran-inspirant.fr/fullhd"

var err error
var DetectedDesktop string

func main() {
	// get os
	DetectedDesktop := os.Getenv("XDG_CURRENT_DESKTOP")
	fmt.Println(DetectedDesktop)

	// "get" current wallpaper to check if we can get it

	// set from url 

	/* if IsGnome(DetectedDesktop) {
		err = setGnomeWallpaperFromFile(ecranInspirantUrl)
	} else {
		err = wallpaper.SetFromURL(ecranInspirantUrl)
	}
	if err != nil {
		panic(err)
	} */

}
