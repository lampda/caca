package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lampda/caca/templates"
	yaml "gopkg.in/yaml.v3"
)

var TMP_PERM = os.ModePerm

func main4() {
	// type uwu struct {
	// 	Name string `yaml:"name"`
	// }
	// u := uwu{
	// 	Name: "\nnya",
	// }
	file := File{
		Name:  "\ntacosdecanaste",
		IsDir: true,
		Files: []*File{
			&File{
				Name: "blah",
				Content: `
func foo()  { fmt.Println("bar");}`,
			},
			&File{}},
	}

	b, err := yaml.Marshal(file)
	panikIfErr(err)
	fmt.Println(string(b))
}

func main() {
	var err error
	input, err := os.Create("input.yaml")
	panikIfErr(err)
	defer input.Close()
	{
		file := File{
			Name:  "tacosdecanaste",
			IsDir: true,
			Files: []*File{
				&File{
					Name:  "cmd",
					IsDir: true,
					Files: []*File{
						&File{
							Name:  "caca2.go",
							IsDir: false,
							Content: `func caca2()  {
								fmt.Println("hello world")
							}`,
						},
						&File{
							Name:  "main.go",
							IsDir: false,
							Content: `func main()  {
								fmt.Println("hello world")
							}`,
						},
						&File{
							Name:  "main.go",
							IsDir: false,
							Content: `func main()  {
								 fmt.Println("hello world")
							}`,
						},
					},
				},
			},
		}

		if env == DEV {
			panikIfErr(os.RemoveAll("tmp"))
		}

		file.Create("tmp")
		// fmt.Println(file.String())

		b, err := yaml.Marshal(file)
		panikIfErr(err)
		assertThisIsTrue(len(b) != 0, "buffer must not be empty")
		n, err := FileWrite(input, b)
		panikIfErr(err)
		assertThisIsTrue(n != 0, "file must not be empty after writing to it")
	}

	{
		b, err := ReadFile(input)
		panikIfErr(err)
		assertThisIsTrue(len(b) != 0, "file must not be empty")
		var fromFile File
		err = yaml.Unmarshal(b, &fromFile)
		panikIfErr(err)
		fmt.Println(fromFile.String())
	}
}

func main2() {
	if len(os.Args) <= 2 {
		crash("u did not provide project name :3 or github user")
	}

	var err error
	projectName := os.Args[1]
	githubUsername := os.Args[2]

	if env == DEV {
		panikIfErr(os.RemoveAll(projectName))
	}

	// TODO: check which permission uses npm/cargo or other project managers for directories and files
	panikIfErr(os.Mkdir(projectName, TMP_PERM))
	panikIfErr(os.Chdir(projectName))

	programPath := fmt.Sprintf("cmd/%s", projectName)
	panikIfErr(os.MkdirAll(programPath, TMP_PERM))

	panikIfErr(os.MkdirAll("bin", TMP_PERM))
	makefile, err := os.Create("Makefile")
	panikIfErr(err)
	defer makefile.Close()
	_, err = os.Create("errors.err")
	panikIfErr(err)
	gitIgnore, err := os.Create(".gitignore")
	panikIfErr(err)
	defer gitIgnore.Close()
	cmd := exec.Command("go", "mod", "init", fmt.Sprintf("github.com/%s/%s", githubUsername, projectName))
	panikIfErr(cmd.Run())

	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", programPath))
	panikIfErr(err)
	defer mainFile.Close()

	loggerFile, err := os.Create(fmt.Sprintf("%s/logger.go", programPath))
	panikIfErr(err)
	defer loggerFile.Close()

	mainFile.WriteString(templates.Main)
	loggerFile.WriteString(templates.Logger)

	makefile.WriteString(fmt.Sprintf(templates.Makefile, projectName))

	fmt.Println("hellope world")
}
