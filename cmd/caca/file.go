package main

import (
	"fmt"
	"os"
	"strings"
)

type Foo struct {
	Bar string
}

type File struct {
	Name    string  `yaml:"name"`
	IsDir   bool    `yaml:"is_dir"`
	Content string  `yaml:"content"`
	Files   []*File `yaml:"files"`
}

func (f *File) StringFormat(sb *strings.Builder) string {
	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("\t\"name\":\"%s\"", f.Name))
	if f.IsDir {
		sb.WriteString(",\t\"files\":[")
		for i, child := range f.Files {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(child.String())
		}
		sb.WriteString("]")
	}
	sb.WriteString("}")
	return sb.String()
}

func (f *File) CreateDirTree(path string) error {
	path = fmt.Sprintf("%s/%s", path, f.Name)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	for _, file := range f.Files {
		err = file.Create(path)
		logErr(err)
		// if err != nil {
		// 	panik("in this case we should do something nasty maybe just logging instead of breaking everything but yeah")
		// }
	}
	return nil
}

func (f *File) Create(path string) error {
	if f.IsDir {
		return f.CreateDirTree(path)
	} else {
		filePath := fmt.Sprintf("%s/%s", path, f.Name)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.WriteString(f.Content)
	}
	return nil
}

func (f *File) String() string {
	var sb strings.Builder
	return f.StringFormat(&sb)
}
