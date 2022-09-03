package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

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
			return fmt.Sprintf("%s:\n%s\n %s",
				pages[i].name,
				strings.Repeat("-", w),
				func() string {
					text, err := ioutil.ReadFile(pages[i].path)
					if err != nil {
						return "ファイルが開けません。"
					}
					return string(text)
				}())
		}))
	if err != nil {
		log.Fatal(err)
	}
	return pages[idx[0]].path
}
