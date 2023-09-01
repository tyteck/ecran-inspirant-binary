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
func TestInArray(t *testing.T) {
	t.Run("do not work with a string", func(t *testing.T) {
		assert.False(t, InArray("Lorem", "Lorem ipsum"))
	})

	t.Run("do not work with empty array", func(t *testing.T) {
		var empty []string
		assert.False(t, InArray("Lorem", empty))
	})

	t.Run("should work with array of string", func(t *testing.T) {
		var values = []string{"Lorem", "ipsum", "dolore", "sit", "amet"}
		assert.True(t, InArray("Lorem", values))
		assert.False(t, InArray("IpSuM", values))
		assert.False(t, InArray("Unknown", values))
	})

	t.Run("should work with array of int", func(t *testing.T) {
		var values = []int{1, 4, 3, 2, 6}
		assert.True(t, InArray(4, values))
		assert.False(t, InArray("4", values))
		assert.False(t, InArray(99, values))
	})

}
