package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Page struct {
	name string
	path string
}

func selectFile() string {
	path := ""
	return path
}

func selectPage(pages []Page) string {
	idx, err := fuzzyfinder.FindMulti(
		pages,
		func(i int) string {
			return pages[i].name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("カテゴリ: %s\n%s\n %s",
				pages[i].name,
        strings.Repeat("-", w),
				func() string {
					// pages[i].path
					text, err := ioutil.ReadFile(filepath.Join(pages[i].path, "README.md"))
					if err != nil {
            return "README.mdが見つかりません。"
					}
					return string(text)
				}())
		}))
	if err != nil {
		log.Fatal(err)
	}
	return pages[idx[0]].path
}
