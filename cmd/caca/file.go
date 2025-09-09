package main

import (
	"fmt"
	"io"
	"os"
)

type File struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

func (f *File) Create(path string) error {
	filePath := fmt.Sprintf("%s/%s", path, f.Name)

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Seek(0, io.SeekStart)

	if f.Content != "" {
		_, err := file.WriteString(f.Content)
		return err
	}

	return nil
}

func (f *File) String() string {
	return f.Name
}
