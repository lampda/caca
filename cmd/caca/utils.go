package main

import (
	"errors"
	"io"
	"os"
)

// TODO: make this return an slice of the actual bytes readed
// like return buffer[:totalBytes]
// and totalBytes+=n each loop
// but i dont remember if the read thing actually puts

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

func FileWriteString(file *os.File, str string) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.WriteString(str)
}

func FileWrite(file *os.File, buffer []byte) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.Write(buffer)
}
