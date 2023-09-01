package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const EcranInspirantUrl = "https://get.ecran-inspirant.fr/fullhd"

// used by gnome, empty every other case
var XdgCurrentDesktop = os.Getenv("XDG_CURRENT_DESKTOP")

var err error

func main() {
	if !IsGnome() {
		fmt.Println("Only gnome is supported for now.")
		os.Exit(1)
	}

	destinationFile := GetFilePath()

	// "get" current wallpaper to check if we can get it
	err = downloadImage(EcranInspirantUrl, destinationFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("download successful, setting new wallpaper ... ")

	SetWallPaperFromFile(destinationFile)

	fmt.Println("Wallpaper has been changed successfully.")
	os.Exit(0)
}

func GetDesktop() (string, error) {
	if IsGnome() {
		return "gnome", nil
	}

	return "", errors.New("not supported yet")
}

func downloadImage(url, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to open url " + url)
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create " + filePath)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("failed to put data " + filePath)
	}

	return nil
}
