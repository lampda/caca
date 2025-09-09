package main

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

var DEV_PERM = os.ModePerm

// create a cool app
type Project struct {
	Name       string `yaml:"name"`
	GithubUser string `yaml:"github_user"`
}

// TODO: handle also creating a git repo with https

func main() {
	if env == DEV {
		panikIfErr(os.RemoveAll("tmp"))
	}

	os.MkdirAll("tmp", DEV_PERM)
	if len(os.Args) <= 2 {
		crash("u did not provide project name :3 or github user")
	}

	panikIfErr(os.RemoveAll(os.Args[1]))

	caca := Caca{
		Name:       os.Args[1],
		GithubUser: os.Args[2],
		Root: &Directory{
			Name: os.Args[1],
			Nodes: []Node{
				&Directory{
					Name: "cmd",
					Nodes: []Node{
						&Directory{
							Name: os.Args[1],
							Nodes: []Node{
								&File{
									Name: "main.go",
									Content: `
package main

func main()  {
	fmt.Println("hola")
}`,
								},
								&File{
									Name: "logger.go",
									Content: `
package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/charmbracelet/lipgloss"
)
func logInfoln(format string, a ...any) {
	logln(INFO, format, a...)
}`,
								},
								&File{
									Name: "helpers.go",
									Content: `
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
}`,
								},
							},
						},
					},
				},
			},
		},
	}

	panikIfErr(caca.Create("./"))

	b, err := yaml.Marshal(caca)
	panikIfErr(err)
	panikIfErr(os.WriteFile("sabrina.yaml", b, fs.ModePerm))
	fmt.Println(string(b))
}
