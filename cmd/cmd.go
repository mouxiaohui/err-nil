package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var PROJECT_DIRECTORY string

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
				if _, err := os.Stat(path); err != nil {
					return err
				}
			}

			PROJECT_DIRECTORY = path

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
