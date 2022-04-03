package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	PROJECT_PATH string
)

func init() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				wd, err := os.Getwd()
				if err != nil {
					return err
				}
				path = wd
			} else {
				if fileInfo, err := os.Stat(path); err != nil {
					return err
				} else {
					if !fileInfo.IsDir() {
						return errors.New(fmt.Sprintf("[%s] 不是文件夹!", path))
					}
				}
			}

			PROJECT_PATH = path

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
