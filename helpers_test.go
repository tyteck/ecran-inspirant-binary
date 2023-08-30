package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveProtocol(t *testing.T) {
	want := "Lorem Ipsum"

	// Define a slice of strings
	prefixes := [5]string{"http://", "ftp://", "https://", "file://", ""}
	for _, prefix := range prefixes {
		got := RemoveProtocol(prefix + want)
		assert.Equal(t, want, got)
	}
}
