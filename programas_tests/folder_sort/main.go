package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb"
)

type FileDefinitions struct {
	Images      []string
	Models      []string
	Sounds      []string
	Videos      []string
	Documents   []string
	Compressed  []string
	Executables []string
	Torrents    []string
	Folders     []string
	Others      []string
	Ignore      []string
}

var files = FileDefinitions{
	Images:      []string{"png", "jpg", "jpeg", "gif", "ico", "psd", "svg"},
	Models:      []string{"stl", "obj"},
	Sounds:      []string{"mp3", "wav", "ogg"},
	Videos:      []string{"mp4", "avi", "mkv", "webm"},
	Documents:   []string{"pdf", "txt", "doc", "docx", "xls", "xlsx", "ppt", "pptx"},
	Compressed:  []string{"zip", "rar", "tar", "gz"},
	Executables: []string{"msi", "win", "exe"},
	Torrents:    []string{"torrent", "bittorent"},
	Others:      []string{"dont_touch"},
	Ignore:      []string{"folder_sort.exe"},
}

var directoryMapping = map[string][]string{
	"./images":      files.Images,
	"./models":      files.Models,
	"./sounds":      files.Sounds,
	"./videos":      files.Videos,
	"./documents":   files.Documents,
	"./compressed":  files.Compressed,
	"./executables": files.Executables,
	"./torrents":    files.Torrents,
	"./others":      files.Others,
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error al obtener la ruta del ejecutable:", err)
		return
	}
	currentDir := filepath.Dir(exePath)
	fmt.Println("La ruta actual de ejecución es:", currentDir)

	dirInfo, _ := os.ReadDir("./")
	var input string
	fmt.Scanln(&input)

	progressBar := pb.StartNew(len(dirInfo))

	fileChan := make(chan os.DirEntry)

	for _, file := range dirInfo {
		go func(file os.DirEntry) {
			fileChan <- file
		}(file)
	}

	for i := 0; i < len(dirInfo); i++ {
		file := <-fileChan
		progressBar.Increment()
		if file.IsDir() {
			if isDirectoryInMapping(file.Name()) {
				continue
			}
			moveFile(file.Name(), true)
			continue
		}
		if contains(files.Ignore, file.Name()) {
			//println("archivo ignorado.")
			continue
		}
		moveFile(file.Name(), false)
	}
	progressBar.Finish()
	fmt.Scanln(&input)
}

func moveFile(fileName string, isFolder bool) {
	extension := filepath.Ext(fileName)
	extension = strings.ReplaceAll(extension, ".", "")

	var destDir string
	for dir, extensions := range directoryMapping {
		if contains(extensions, extension) {
			destDir = dir
			break
		}
	}

	if destDir == "" {
		if isFolder {
			destDir = "./folders"
		} else {
			destDir = "./others"
		}
	}
	os.Mkdir(destDir, 0700)
	newPath := filepath.Join(destDir, fileName)
	originalPath := filepath.Join("./", fileName)
	err := os.Rename(originalPath, newPath)
	if err != nil {
		fmt.Println("Error al mover el archivo:", err)
	} else {
		//fmt.Printf("Archivo %s movido a %s\n", fileName, newPath)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isDirectoryInMapping(dirName string) bool {
	for destDir := range directoryMapping {
		if strings.TrimPrefix(destDir, "./") == dirName {
			return true
		}
	}
	return false
}
