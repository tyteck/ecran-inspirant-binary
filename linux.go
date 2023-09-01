package main

import (
	"os"
)

func GetFilePath() string {
	return os.TempDir() + "/" + YoutubeId() + ".jpg"
}
