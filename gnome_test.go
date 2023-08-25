package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	want := "default"
	got := defaultTheme

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIsGnomeReturnFalse(t *testing.T) {
	var got bool

	// Define an array of strings
	validStrings := [3]string{"Lorem", "Windows", "gnome"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = IsGnome(stringToTest)
		assert.False(t, got)
	}
}

func TestIsGnomeReturnTrue(t *testing.T) {
	var got bool

	// Define an array of strings
	validStrings := [4]string{"GNOME", "Unity", "Pantheon", "Lorem ipsum GNOME"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = IsGnome(stringToTest)
		assert.True(t, got)
	}
}

func TestCurrentGnomeTheme(t *testing.T) {
	var want string
	got, _ := CurrentGnomeTheme()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
