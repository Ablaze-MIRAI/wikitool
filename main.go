package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "new",
				Usage: "新しいページを作成します。",
				Action: func(*cli.Context) error {
					// プロンプトを用いて対話的にページを作成する

					categoryPath := selectCategory(categoryList())
					fileName := inputFileName()
					newFilePath := replaceExt(
						filepath.Join(categoryPath, fileName),
						"",
						".md")
					Xopen(newFilePath)
					return nil
				},
			},
			{
				Name:  "edit",
				Usage: "ページを選択して編集します。",
				Action: func(*cli.Context) error {
					path := selectCategory(categoryList())
          // fmt.Println("Filepath", path)
          // fmt.Println("pageList", pageList(path))
          base := selectPage(pageList(path))
          fmt.Println(base)
					Xopen(base)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
