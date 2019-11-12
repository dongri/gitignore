package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := findAll("gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		fmt.Println(f)
	}
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
