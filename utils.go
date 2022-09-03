package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"fmt"
	"github.com/manifoldco/promptui"
)

type Category struct {
	name string
	path string
}

type Page struct {
  name string
  path string
}

func Xopen(path string) {
	var command string

	if runtime.GOOS == "windows" {
		command = "start"

		open(command, path)
	} else if runtime.GOOS == "linux" {
		visual := os.Getenv("VISUAL")
		editor := os.Getenv("EDITOR")

		if visual != "" {
			command = visual
			open(command, path)
		} else if editor != "" {
			command = editor
			open(command, path)
		} else {
			command = "/usr/bin/edit"
			open(command, path)
		}
	}
}

func open(cmd, path string) {
	c := exec.Command(cmd, path)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

func categoryList() []Category {
	// ページ一覧を取得

	var path string
	var categories []Category

	cwd, _ := os.Getwd()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == ".git" || !file.IsDir() {
			continue
		}

		path = filepath.Join(cwd, file.Name())

		categories = append(categories, Category{
			name: file.Name(),
			path: path,
		})
	}
	return categories
}

func genNewPath(categoryPath, fileName string) string {
	return filepath.Join(categoryPath, fileName)
}

func inputFileName() string {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "ファイル名",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func pageList(searchDir string) []Page {
  var pages []Page

  cwd, _ := os.Getwd()

	files, err := ioutil.ReadDir(searchDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == ".git" || file.Name() == "imgs" || file.IsDir() {
			continue
		}

    filePath := filepath.Join(cwd, file.Name())
    fmt.Println(file.Name())

    pages = append(pages, Page{
      name: file.Name(),
      path: filePath,
    })
	}

  return pages
}
