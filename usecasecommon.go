package main

import (
	"os"
	"path/filepath"
	"strings"
)

func strsToPath(args ...string) (resp string) {

	for i, v := range args {
		resp = resp + v
		if i < len(args)-1 {
			resp = resp + "/"
		}
	}

	return resp
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getGenre(fileName string) string {
	var (
		genre string
	)

	lastIdx := strings.LastIndex(fileName, separator)
	genre = fileName[:lastIdx]
	genre = fileName[lastIdx+1:]

	genre = strings.TrimSuffix(genre, filepath.Ext(extPNG))

	return genre
}

func getDramaName(fileName string) string {
	var (
		name string
	)

	name = fileName[:strings.IndexByte(fileName, separatorByte)]

	return name

}
