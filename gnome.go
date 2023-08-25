package main

import (
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

const defaultTheme string = "default"
const darkTheme string = "prefer-dark"

// const lightTheme = "prefer-light"

func IsGnome(s string) bool {
	return strings.Contains(s, "GNOME") || s == "Unity" || s == "Pantheon"
}

func CurrentGnomeWallpaper() (string, error) {
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
	if !IsGnome(DetectedDesktop) {
		return "", nil
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
