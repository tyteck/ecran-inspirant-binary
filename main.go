package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/reujab/wallpaper"
)

const ecranInspirantUrl = "get.ecran-inspirant.fr/fullhd"

var err error
var DetectedDesktop string

func main() {
	// get os
	DetectedDesktop := os.Getenv("XDG_CURRENT_DESKTOP")
	fmt.Println(DetectedDesktop)

	// "get" current wallpaper to check if we can get it
	//err = downloadImage(ecranInspirantUrl)

	// set from url

	if IsGnome(DetectedDesktop) {
		//err = SetWallpaperFromUrl(ecranInspirantUrl)
		err = fmt.Errorf("to be continued")
	} else {
		err = wallpaper.SetFromURL(ecranInspirantUrl)
		err = fmt.Errorf("to be continued")
	}
	if err != nil {
		panic(err)
	}
}

func downloadImage(url, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}
