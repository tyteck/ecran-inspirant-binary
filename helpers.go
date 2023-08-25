package main

import "strings"

func RemoveProtocol(s string) string {
	s = strings.TrimPrefix(s, "http://")
	s = strings.TrimPrefix(s, "https://")
	s = strings.TrimPrefix(s, "file://")
	return s
}
