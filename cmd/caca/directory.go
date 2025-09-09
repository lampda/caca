package main

import (
	"fmt"
	"os"
)

type Directory struct {
	Name  string `yaml:"name"`
	Nodes []Node `yaml:"nodes"`
}

func (dir *Directory) Create(path string) error {
	path = fmt.Sprintf("%s/%s", path, dir.Name)
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	for _, file := range dir.Nodes {
		err = file.Create(path)
		logErr(err)
		// if err != nil {
		// 	panik("in this case we should do something nasty maybe just logging instead of breaking everything but yeah")
		// }
	}

	return nil
}
