// go:+build linux
package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const defaultTheme string = "default"
const darkTheme string = "prefer-dark"
const lightTheme = "prefer-light"

func IsGnome() bool {
	return strings.Contains(XdgCurrentDesktop, "GNOME") || XdgCurrentDesktop == "Unity" || XdgCurrentDesktop == "Pantheon"
}

func CurrentGnomeWallPaper() (string, error) {
	currentTheme, err := CurrentGnomeTheme()
	if err != nil {
		return "", err
	}
	switch currentTheme {
	case darkTheme:
		return parseDconf("gsettings", "get", "org.gnome.desktop.background", "picture-uri-dark")
	default:
		return parseDconf("gsettings", "get", "org.gnome.desktop.background", "picture-uri")
	}
}

func CurrentGnomeTheme() (string, error) {
	if !IsGnome() {
		return "", fmt.Errorf("not a gnome desktop, you should not use it")
	}
	return parseDconf("gsettings", "get", "org.gnome.desktop.interface", "color-scheme")
}

func parseDconf(command string, args ...string) (string, error) {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}

	// unquote string
	var unquoted string

	// the output is quoted with single quotes, which cannot be unquoted using strconv.Unquote, but it is valid yaml
	err = yaml.UnmarshalStrict(output, &unquoted)
	if err != nil {
		return unquoted, err
	}

	return RemoveProtocol(unquoted), nil
}

func SetWallPaperFromFile(filePath string) error {
	currentTheme, err := CurrentGnomeTheme()
	if err != nil {
		return err
	}

	var keyToUse string
	switch currentTheme {
	case darkTheme:
		keyToUse = "picture-uri-dark"
	default:
		keyToUse = "picture-uri"
	}

	return exec.Command("gsettings", "set", "org.gnome.desktop.background", keyToUse, strconv.Quote("file://"+filePath)).Run()

}
