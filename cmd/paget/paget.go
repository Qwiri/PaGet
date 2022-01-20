package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/Qwiri/PaGet/internal/colors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Name = "PaGet"
	app.Usage = "Simple, opionionated package manager"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		fmt.Println(colors.Red + "Command not implemented yet 😭" + colors.Reset)
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install one or mutiple packages",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "Install one or mutiple packages",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "upgrade",
			Usage: "Install one or mutiple packages",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
