package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/manifoldco/promptui"
)

func selectCategory(categories []Category) string {
	cwd, _ := os.Getwd()

	idx, err := fuzzyfinder.FindMulti(
		categories,
		func(i int) string {
			return categories[i].name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("カテゴリー: %s\n%s\n%s",
				categories[i].name,
				strings.Repeat("-", w),
				func() string {
					// pages[i].path
					text, err := ioutil.ReadFile(
						filepath.Join(
							cwd,
							categories[i].name,
							"README.md"))

					if err != nil {
						return "カテゴリーの説明がありません。"
					}
					return string(text)
				}())
		}))
	if err != nil {
		log.Fatal(err)
	}
	return categories[idx[0]].path
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

  fmt.Printf("You choose %q\n", result)
  return result
}
