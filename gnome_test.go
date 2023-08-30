package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	want := "default"
	got := defaultTheme

	assert.Equal(t, want, got)
}

func TestIsGnomeReturnFalse(t *testing.T) {
	var got bool

	// Define a slice of strings
	validStrings := []string{"Lorem", "Windows", "gnome"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = IsGnome(stringToTest)
		assert.False(t, got)
	}
}

func TestIsGnomeReturnTrue(t *testing.T) {
	var got bool

	// Define a slice of strings
	validStrings := []string{"GNOME", "Unity", "Pantheon", "Lorem ipsum GNOME"}

	// Loop over the array elements using range
	for _, stringToTest := range validStrings {
		got = IsGnome(stringToTest)
		assert.True(t, got)
	}
}
