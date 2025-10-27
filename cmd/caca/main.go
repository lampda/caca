package caca

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"
)

var (
	ErrIsNotValidTextFile = fmt.Errorf("file is not readable")
)

func readFileIfIsText(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	if fileScanner.Err() != nil {
		return nil, err
	}

	if !utf8.ValidString(fileScanner.Text()) {
		return nil, nil
	}

	b := fileScanner.Bytes()

	for {
		keep := fileScanner.Scan()
		if fileScanner.Err() != nil {
			return nil, err
		}
		if !keep {
			break
		}

		b = append(b, fileScanner.Bytes()...)
		b = append(b, '\n')
	}

	return b, nil
}

type CACA struct {
	projectName           string
	originalProjectName   string
	files                 []File
	filesContentToReplace []string
}

type File struct {
	fullPath string
	path     string
	name     string
	content  string
	mode     fs.FileMode
}

func (c *CACA) analizeDir(path string, d fs.DirEntry, err error) error {

	if d.IsDir() {
		return nil
	}
	info, err := d.Info()

	if err != nil {
		return err
	}

	b, err := readFileIfIsText(path)

	if err != nil {
		return err
	}

	content := string(b)

	/*
		if the curr file is a binary file
		is not an error, but there is not
		content, and we are just copying
		the files that are text
	*/

	if b != nil {
		// match := fmt.Sprintf(".*/%s", c.originalBasePath)
		// regexp.Match(match, path)

		newPath := sillySwapProjectName(path, c.originalProjectName, c.projectName)
		currFileName := d.Name()

		if slices.Contains(c.filesContentToReplace, currFileName) {
			content = strings.ReplaceAll(content, c.originalProjectName, c.projectName)
		}

		c.files = append(c.files, File{
			fullPath: newPath,
			path:     filepath.Dir(newPath),
			name:     currFileName,
			content:  content,
			mode:     info.Mode(),
		})

	}

	return nil
}

var (
	EMPTY = "<empty>"
)

type dirRegex struct {
	FileName string `json:"file_name"`
	Regex    string `json:"regex"`
	Replace  string `json:"replace"`
}

type template struct {
	TemplateName          string     `json:"template_name"`
	Path                  string     `json:"path"`
	FilesContentToReplace []string   `json:"files_to_replace"`
	Regexes               []dirRegex `json:"regexes"`
}

type config struct {
	DefaultTemplate string     `json:"default_template"`
	Templates       []template `json:"templates"`
}

func main() {

	var projectName string
	var templatePath string
	var templateName string
	var filesContentToReplace []string

	flag.StringVar(&projectName, "name", EMPTY, "Name for the project")
	flag.StringVar(&templatePath, "template-path", EMPTY, "Template's path")
	flag.StringVar(&templateName, "template-name", EMPTY, "Template name to use")
	home := os.Getenv("HOME")
	flag.Parse()

	if projectName == EMPTY {
		logError("please provide a project name")
		os.Exit(-1)
	}

	if templatePath == EMPTY {
		b, err := os.ReadFile(fmt.Sprintf("%s/.config/caca/caca.json", home))

		if err != nil {
			prettyLogErr(err, "please provide a template path")
			os.Exit(-1)
		}

		var cfg config
		err = json.Unmarshal(b, &cfg)

		if err != nil {
			logErr(err)
			os.Exit(-1)
		}

		targetTemplateName := cfg.DefaultTemplate

		if templateName != EMPTY {
			targetTemplateName = templateName
		}

		for _, t := range cfg.Templates {
			if t.TemplateName == targetTemplateName {
				templatePath = t.Path
				filesContentToReplace = t.FilesContentToReplace
				break
			}
		}

	}

	var err error

	// this thing replaces
	b := filepath.Base(templatePath)
	caca := CACA{
		projectName:           projectName,
		files:                 nil,
		originalProjectName:   b,
		filesContentToReplace: slices.Concat(filesContentToReplace, []string{"Makefile"}),
	}

	err = filepath.WalkDir(templatePath, caca.analizeDir)

	if err != nil {
		logErr(err)
		os.Exit(-1)
	}

	for _, f := range caca.files {
		err := os.MkdirAll(f.path, os.ModePerm)
		if err != nil {
			logErr(err)
			continue
		}
		err = os.WriteFile(f.fullPath, []byte(f.content), f.mode)
		if err != nil {
			logErr(err)
			continue
		}
	}

}
