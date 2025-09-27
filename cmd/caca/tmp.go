package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type node_t interface {
	Create(path string) error
}

type Silly struct {
	Path string `json:"path"`
}

type dir_t struct {
	// Name  string   `json:"name"`
	// Silly Silly    `json:"silly"`
	Files []node_t `json:"files"`
}

func (d dir_t) Create(path string) error {
	fmt.Println("create for dir", path)
	return nil
}

type file_t struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (f file_t) Create(path string) error {
	fmt.Println("create for file", path)
	return nil
}

type root_t struct {
	Uwu    string `json:"uwu"`
	Parent node_t `json:"parent"`
}

func (r *root_t) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	parent := raw["parent"].(map[string]interface{})
	if parent["files"] != nil {
		var d dir_t
		convertMapToStructure(parent, &d)
		fmt.Println("dir:")
		fmt.Println(d)
	} else {
		var f file_t
		convertMapToStructure(parent, &f)
		fmt.Println("file:")
		fmt.Println(f)
	}
	return nil
}

// stolen from: https://www.golinuxcloud.com/go-map-to-struct/
// check mitchellh thing: github.com/mitchellh/mapstructure

func main() {
	caca()
	// initCaca()
	// dirtest()
}

func dirtest2() {
	jsonNode := map[string]any{
		"name": "supercoolproject",
		"songs": []map[string]any{
			map[string]any{
				"lyrics": "iuiuasidf",
			},
		},
	}
	var sk Silk
	panikIfErr(convertMapToStructure(jsonNode, &sk))
	fmt.Println(sk)
	// readJson()
}

func dirtest() {
	jsonNode := map[string]any{
		// "name": "supercoolproject",
		// "silly": map[string]any{
		// 	"path": "/var/db/Makefile",
		// },
		"files": []map[string]any{
			map[string]any{
				"name":    "owo.go",
				"content": "fmt.Println('hello');",
			},
		},
	}
	var n dir_t
	panikIfErr(convertMapToStructure(jsonNode, &n))
	fmt.Println("hellope", n.Files)
	// readJson()
}

func smalltest() {
	jsonNode := map[string]interface{}{
		"name":    "owo.go",
		"content": "fmt.Println('hello');",
	}
	var n file_t
	panikIfErr(convertMapToStructure(jsonNode, &n))
	fmt.Println(n)
}

func interfacing() {
	f := map[string]interface{}{
		"name": "foo",
		"files": []map[string]interface{}{
			map[string]interface{}{
				"name":    "owo.go",
				"content": "fmt.Println('hello');",
			},
		},
	}
	fmt.Println(f)
}

func readJson() {
	var input root_t
	b, err := os.ReadFile("example.json")
	panikIfErr(err)
	err = json.Unmarshal(b, &input)
	panikIfErr(err)
	fmt.Println(input)
}

func writeJson() {
	output := root_t{
		Uwu: "owo",
		Parent: dir_t{
			// Name: "foo",
			Files: []node_t{
				file_t{
					Name:    "owo.go",
					Content: "fmt.Println(\"nya\");",
				},
			},
		},
	}
	b, err := json.Marshal(output)
	panikIfErr(err)
	os.WriteFile("example.json", b, DEV_PERM)
}
