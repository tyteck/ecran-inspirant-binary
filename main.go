package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/reujab/wallpaper"
)

const ecranInspirantUrl = "get.ecran-inspirant.fr/fullhd"

var err error
var DetectedDesktop string

func main() {
	// get os
	DetectedDesktop := os.Getenv("XDG_CURRENT_DESKTOP")

	if IsGnome(DetectedDesktop) {
		err = setGnomeWallpaperFromFIle(ecranInspirantUrl)
	} else {
		err = wallpaper.SetFromURL(ecranInspirantUrl)
	}
	if err != nil {
		panic(err)
	}

}

func setGnomeWallpaperFromFIle(file string) error {
	switch theme, _ := CurrentGnomeTheme(); theme {
	case darkTheme:
		println(strconv.Quote("file://" + file))
		return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", strconv.Quote("file://"+file)).Run()
	default:
		return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+file)).Run()
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
