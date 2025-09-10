// TODO: help:33

package main

import (
	"os"

	"github.com/ghodss/yaml"
)

type node interface {
	Create(string)
}

type file struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type directory struct {
	Name  string `json:"name"`
	Nodes []node `json:"nodes"`
}

func (f *file) Create(s string)      {}
func (d *directory) Create(s string) {}

type root struct {
	ProjectName string `json:"project_name"`
	Parent      node   `json:"parent"`
}

func interfaceDebugUnmarshal() {
	var r root
	b, err := os.ReadFile("/home/marcig/personal/summer/caca/test/dolls.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
}
