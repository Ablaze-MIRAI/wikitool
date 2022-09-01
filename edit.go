package main

import (
	"fmt"
	"io/ioutil"
	"log"

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
			return fmt.Sprintf("Track: %s (%s)\n",
				pages[i].name,
				func() string {
					// pages[i].path
					text, err := ioutil.ReadFile(pages[i].path)
					if err != nil {
						log.Fatal(err)
					}
					return string(text)
				}())
		}))
	if err != nil {
		log.Fatal(err)
	}
	return pages[idx[0]].path
}
