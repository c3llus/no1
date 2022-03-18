package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	username      = os.Getenv("USER")
	baseDirectory = fmt.Sprintf(baseDirectoryFormat, username)
	filePathZip   = strsToPath(basePath, zipFileName)
)

func CreateBaseDirectory() {
	baseDirectory = fmt.Sprintf(baseDirectoryFormat, username)
	os.MkdirAll(baseDirectory, os.ModePerm)
}

// JAWABAN A
func ExtractZip() {
	archive, err := zip.OpenReader(filePathZip)
	if err != nil {
		return
	}
	defer archive.Close()

	for _, f := range archive.File {

		// Exclude non-drakor
		if isNotDrakor(f) {
			continue
		}

		filePath := DrakorSorter(f.FileInfo().Name())
		fmt.Println("PATH: " + filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(baseDirectory)+string(os.PathSeparator)) {
			return
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func isNotDrakor(f *zip.File) bool {
	// Dont extract folder
	if f.FileInfo().IsDir() {
		return true
	}

	// Dont extract non-PNG
	if filepath.Ext(f.FileInfo().Name()) != extPNG {
		return true
	}

	return false
}

// JAWABAN B
func DrakorSorter(fileName string) (fileDest string) {

	var (
		filePath string
		genre    string
	)

	genre = getGenre(fileName)
	filePath = strsToPath(baseDirectory, genre)

	// If dir not exist, create
	if !pathExists(filePath) {
		os.MkdirAll(filePath, os.ModePerm)
	}

	// JAWABAN C
	fileName = getDramaName(fileName) + extPNG

	return strsToPath(baseDirectory, genre, fileName)

}
