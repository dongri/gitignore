package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/c-bata/go-prompt"
)

var files []string

func main() {
	var err error
	gopath := os.Getenv("GOPATH")
	files, err = findAll(gopath + "/src/github.com/dongri/gitignore/github/gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in := prompt.Input("Search: ", completer)

	var targetFile string
	for _, f := range files {
		file := lastString(strings.Split(f, "/"))
		if file == in {
			targetFile = f
			break
		}
	}
	fmt.Println("Generated: " + in)
	err = exec.Command("cp", targetFile, "./.gitignore").Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	var s = []prompt.Suggest{}
	for _, f := range files {
		file := lastString(strings.Split(f, "/"))
		suggest := prompt.Suggest{
			Text:        file,
			Description: file,
		}
		s = append(s, suggest)
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func findAll(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	var paths []string
	for _, file := range files {
		if !file.IsDir() && !strings.HasSuffix(file.Name(), ".gitignore") {
			continue
		}
		if file.IsDir() {
			files, err := findAll(filepath.Join(dir, file.Name()))
			if err != nil {
				return []string{}, err
			}
			paths = append(paths, files...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}
