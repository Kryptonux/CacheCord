package util

import (
	"path/filepath"
	"strings"
)

func ParseURL(urlString string) (string, string) {
	u := strings.TrimRight(urlString, "/")
	parts := strings.Split(u, "/")
	filename := parts[len(parts)-1]
	format := filepath.Ext(filename)[1:]
	return filename, format
}
