//go:build windows

package main

import (
	"os"
)

func getCacheDir() (string, error) {
	return os.TempDir(), nil
}
