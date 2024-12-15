package find

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DirectoryContent struct {
	folders     []string
	files       []string
	symbolLinks map[string]string
}

func (dc *DirectoryContent) Folders() []string {
	return dc.folders
}

func (dc *DirectoryContent) SymbolLinks() map[string]string {
	return dc.symbolLinks
}

func (dc *DirectoryContent) Files() []string {
	return dc.files
}

func ReadDir(filePath string, dc *DirectoryContent) error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			log.Printf("Permission denied: %s", filePath)
			return nil
		}
		return fmt.Errorf("failed to read directory %s: %w", filePath, err)
	}
	for _, file := range files {
		filename := filepath.Join(filePath, file.Name())
		switch {
		case file.IsDir():
			dc.folders = append(dc.folders, filename)
		case isLink(filename):
			dc.addLink(filename)
		default:
			dc.files = append(dc.files, filename)
		}
	}
	for _, file := range files {
		if file.IsDir() {
			ReadDir(filepath.Join(filePath, file.Name()), dc)
		}
	}
	return nil
}

func ReadContent(filepath string) (DirectoryContent, error) {
	dc := DirectoryContent{
		folders:     []string{},
		files:       []string{},
		symbolLinks: map[string]string{},
	}

	err := ReadDir(filepath, &dc)

	return dc, err
}

func (dc *DirectoryContent) addLink(filename string) {

	targetPath, err := filepath.EvalSymlinks(filename)
	if err != nil {
		dc.symbolLinks[filename] = "[broken]"
	} else {
		dc.symbolLinks[filename] = targetPath
	}
}

func isLink(filename string) bool {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}
	isLink := fileInfo.Mode()&os.ModeSymlink != 0
	return isLink
}

func (dc *DirectoryContent) PrintFolders(out io.Writer) {
	for _, file := range dc.Folders() {
		fmt.Fprintln(out, file)
	}
}

func (dc *DirectoryContent) PrintFiles(out io.Writer, extension string) {
	for _, file := range dc.Files() {
		if strings.HasSuffix(file, extension) {
			fmt.Fprintln(out, file)
		}
	}
}

func (dc *DirectoryContent) PrintSymbolLinks(out io.Writer) {
	for k, v := range dc.SymbolLinks() {
		fmt.Fprintln(out, k, "->", v)
	}
}
