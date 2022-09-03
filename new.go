package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
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
						return "ディレクトリ内にREADME.mdが見つかりません。"
					}
					return string(text)
				}())
		}))

	if err != nil {
		log.Fatal(err)
	}
	return categories[idx[0]].path
}
