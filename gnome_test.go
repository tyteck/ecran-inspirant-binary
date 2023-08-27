package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	want := "default"
	got := defaultTheme

	assert.Equal(t, got, want)
}

func TestIsGnomeReturnFalse(t *testing.T) {
	var got bool

	// Define a slice of strings
	validStrings := [3]string{"Lorem", "Windows", "gnome"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = isGnome(stringToTest)
		assert.False(t, got)
	}
}

func TestIsGnomeReturnTrue(t *testing.T) {
	var got bool

	// Define a slice of strings
	validStrings := [4]string{"GNOME", "Unity", "Pantheon", "Lorem ipsum GNOME"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = isGnome(stringToTest)
		assert.True(t, got)
	}
}
