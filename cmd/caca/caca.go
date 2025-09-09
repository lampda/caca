package main

import (
	"fmt"
	"os/exec"
)

type Node interface {
	Create(path string) error
}

type Root struct {
	name string
	root Node
}

type Caca struct {
	Name       string   `yaml:"name"`
	GithubUser string   `yaml:"github_user"`
	Root       Node     `yaml:"root"`
	Packages   []string `yaml:"packages"`
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
