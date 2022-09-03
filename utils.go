package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Category struct {
	name string
	path string
}

func pageList() []Page {
	// ページ一覧を取得

	var path string
	var pages []Page

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
    fmt.Println(path)

		pages = append(pages, Page{
			name: file.Name(),
			path: path,
		})
	}
	return pages
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

func categorys() []Category {
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

func replaceExt(filePath, from, to string) string {
    ext := filepath.Ext(filePath)
    if len(from) > 0 && ext != from {
        return filePath
    }
    return filePath[:len(filePath)-len(ext)] + to
}
