package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate += "\nCUSTOMIZED: you bet ur muffins\n"
	cli.CommandHelpTemplate += "\nYMMV\n"
	cli.SubcommandHelpTemplate += "\nor something\n"

	cli.HelpFlag = &cli.BoolFlag{Name: "help", Aliases: []string{"h"}}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"v"}}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Println(templ)
		fmt.Fprintf(w, "best of luck to you\n")
	}
	cli.VersionPrinter = func(c *cli.Context) {
		out := color.New(color.FgBlue).Add(color.Bold)
		out.Fprintf(c.App.Writer, "Version %s\n", c.App.Version)
	}
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}
	cli.ErrWriter = ioutil.Discard
	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.Names()[0])
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "PaGet"
	app.Usage = "Simple, opionionated package manager"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "Marc Türke",
			Email: "email",
		},
	}

	app.Action = func(c *cli.Context) error {
		out := color.New(color.FgRed).Add(color.Bold)
		out.Println("Not implemented yet!")
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
