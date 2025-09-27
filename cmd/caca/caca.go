package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// type Node interface {
// 	Create(path string) error
// 	// UnmarshalJSON() ([]byte, error)
// }

type Node struct {
	IsDir   bool   `json:"is_dir"`
	Name    string `json:"name"`
	Files   []Node `json:"files"`
	Content string `json:"content"`
}

type Caca struct {
	Name       string   `json:"name"`
	GithubUser string   `json:"github_user"`
	Root       Node     `json:"root"`
	Packages   []string `json:"packages"`
}

// TODO: throw a warning in case files share the same name in same directory
func (caca *Caca) Create(path string) error {
	return caca.Root.Create(path)
}
func (caca *Caca) GoModInit() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	os.Chdir(caca.Name)
	githubRepo := fmt.Sprintf("github.com/%s/%s", caca.GithubUser, caca.Name)
	cmd := exec.Command("go", "mod", "init", githubRepo)
	// TODO: handle errors here
	err = cmd.Run()
	logErr(err)
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	logErr(err)
	os.Chdir(wd)
	return nil
}

func (node *Node) Create(path string) error {
	if len(node.Files) > 0 {
		path = fmt.Sprintf("%s/%s", path, node.Name)
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			return err
		}

		for _, file := range node.Files {
			err = file.Create(path)
			logErr(err)
		}
		return nil

	} else {
		filePath := fmt.Sprintf("%s/%s", path, node.Name)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Seek(0, io.SeekStart)
		if node.Content != "" {
			_, err := file.WriteString(node.Content)
			return err
		}
		return nil

	}
	return nil
}

func (caca *Caca) InitGit() {
	cmd := exec.Command("git", "init")
	githubRepo := fmt.Sprintf("git@github.com:%s/%s:git", caca.GithubUser, caca.Name)
	panikIfErr(cmd.Run())
	cmd = exec.Command("git", "branch", "-M", "main")
	panikIfErr(cmd.Run())
	cmd = exec.Command(
		"git",
		"remote",
		"add",
		"origin",
		githubRepo,
	)
	panikIfErr(cmd.Run())
}
