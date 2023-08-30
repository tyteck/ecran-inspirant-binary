package main

import (
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCacheDir(t *testing.T) {
	currentUser, _ := user.Current()
	got, _ := getCacheDir()

	want := filepath.Join(currentUser.HomeDir, ".cache")
	assert.Equal(t, want, got)
}
