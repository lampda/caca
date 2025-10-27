package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

// TODO: make this return an slice of the actual bytes readed
// like return buffer[:totalBytes]
// and totalBytes+=n each loop
// but i dont remember if the read thing actually puts

// https://stackoverflow.com/questions/64925992/write-and-read-file-with-same-os-file-in-go
func ReadFile(file *os.File) ([]byte, error) {
	defer file.Seek(0, io.SeekStart)
	buffer := make([]byte, 1024)
	var totalBytes int
	for {
		n, err := file.Read(buffer)
		totalBytes += n
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				return buffer[:totalBytes], nil
			default:
				return nil, err
			}
		}
	}
}

func cleanPath(path, parent string) string {
	pathSplitted := splitPath(path)

	for i, name := range pathSplitted {
		if name == parent {
			return strings.Join(pathSplitted[i:], string(os.PathSeparator))
		}
	}

	return ""
}

// NOTE: this thing replaces all the files with the originalParent substring with the newParent
// so this might be actually not what the user wants, cuz maybe there is a file with the name of the project, like, ebit_utils.go which is for things to work with ebit, and if we copy the template well, silly things could happen, but at this point i dont care that much!
func sillySwapProjectName(path, oldParent, newParent string) string {
	relativePath := cleanPath(path, oldParent)
	return strings.ReplaceAll(relativePath, oldParent, newParent)
}

func splitPath(path string) []string {
	if path == "" {
		return []string{}
	}
	return strings.Split(path, string(os.PathSeparator))
}

func FileWriteString(file *os.File, str string) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.WriteString(str)
}

func FileWrite(file *os.File, buffer []byte) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.Write(buffer)
}
