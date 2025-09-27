
package main

import (
	"errors"
	"io"
	"os"
)

func FileWriteString(file *os.File, str string) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.WriteString(str)
}

func FileWrite(file *os.File, buffer []byte) (int, error) {
	defer file.Seek(0, io.SeekStart)
	return file.Write(buffer)
}