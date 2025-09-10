package main

import (
	"fmt"
	"os/exec"
)

type Node interface {
	Create(path string) error
	// UnmarshalJSON() ([]byte, error)
}

// type Node struct {
// 	IsDir   bool    `json:"is_dir"`
// 	Name    string  `json:"name"`
// 	Files   []*Node `json:"files"`
// 	Content string  `json:"content"`
// }

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
