package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"

	"github.com/ghodss/yaml"
)

var DEV_PERM = os.ModePerm

// create a cool app

type Fuzz interface {
	fizz() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type Silk struct {
	Hornet bool `json:"hornet"`
}

type Song struct {
	Bababui string `json:"bababui"`
}

func (s *Song) fizz() string {
	return "fizz"
}

func (s *Song) MarshalJSON() ([]byte, error) {
	fmt.Println("marshal")
	return []byte(s.Bababui), nil
}

func (s *Song) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshal")
	s.Bababui = string(data)
	return nil
}

func (s *Silk) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(s.Hornet)), nil
}

func (s *Silk) UnmarshalJSON(data []byte) error {
	hornet, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	s.Hornet = hornet
	return nil
}

func (s *Silk) fizz() string {
	return "fizz"
}

// func (s Silk) UnmarshalJSON(buffer []byte, output interface{}) error {
// 	fmt.Println(string(buffer))
// 	return nil
// }

type Project struct {
	Name        string `json:"name"`
	GithubUser  string `json:"github_user"`
	ProjectNode Node   `json:"project_node"`
}

type Foo struct {
	Doe          string   `json:"doe"`
	Ray          string   `json:"ray"`
	Pi           float32  `json:"pi"`
	Xmas         bool     `json:"xmas"`
	FrenchHens   int      `json:"french-hens"`
	CallingBirds []string `json:"calling-birds"`
	XmasFifthDay struct {
		callingBirds string `json:"calling-birds"`
		frenchHens   int    `json:"french-hens"`
		goldenRings  int    `json:"golden-rings"`
		Partridges   struct {
			CallingBirds string `json:"calling-birds"`
			FrenchHens   int    `json:"french-hens"`
			GoldenRings  int    `json:"golden-rings"`
		} `json:"partridges"`
		TurtleDoves int `json:"turtle-doves"`
	} `json:"xmas-fifth-day"`
}

// TODO: handle also creating a git repo with https

func main() {
	// interfaceDebugUnmarshal()
	// nodes := []map[string]interface{}{
	// 	map[string]interface{}{
	// 		"content": "func help(){}",
	// 		"name":    "helpers.go",
	// 	},
	// 	map[string]interface{}{
	// 		"content": "func main() {fmt.Println(hellope);}",
	// 		"name":    "main.go"},
	// }
	// for _, n := range nodes {
	// 	AssertFile(n)
	// }
	// interfaceSilly()
	// interfaceSilly()
	// interface2()
	// interfaceDebug()
	write()
	read()
	// if false {
	// } else {
	// }
}

func read() {
	if env == DEV {
		panikIfErr(os.RemoveAll("tmp"))
	}

	os.MkdirAll("tmp", DEV_PERM)

	if len(os.Args) <= 2 {
		crash("u did not provide project name :3 or github user")
	}

	panikIfErr(os.RemoveAll(os.Args[1]))
	var prj Project
	buffer, err := os.ReadFile("/home/marcig/personal/summer/caca/test/sabrina.yaml")
	panikIfErr(err)
	err = yaml.Unmarshal(buffer, &prj)
	panikIfErr(err)
	fmt.Println(string(buffer))
	// fmt.Println(prj)
}

func write() {
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
	// fmt.Println(string(b))
}
