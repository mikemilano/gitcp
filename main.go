package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "seeder"
	app.Usage = "copy source from remote repository"
	app.Version = "0.0.1"

	// TODO: Add option for ssh key
	// TODO: Add option for github https vs ssh

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "src, s",
			Usage: "source path to copy from; separate multiple with commas",
			Value: "./",
		},
		cli.StringFlag{
			Name:  "dst, d",
			Usage: "destination path to copy to; separate multiple with commas",
			Value: "./",
		},
		cli.StringFlag{
			Name:     "clone-cdir, c",
			Usage:    "temp directory where repo will be cloned; \"memory\" to use system memory;",
			FilePath: "/tmp",
			Value: "/tmp",
		},
		cli.StringFlag{
			Name:  "github-proto, p",
			Usage: "Github proto; auto, https, or ssh",
			Value: "auto",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "quiet mode, no output",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "verbose mode for troubleshooting",
		},
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "print the version",
	}

	app.Action = func(c *cli.Context) error {
		target := c.Args().Get(0)

		ci := ConfigInput{}

		if c.String("src") != "" {
			ci.src = c.String("src")
		}
		if c.String("dst") != "" {
			ci.dst = c.String("dst")
		}
		if c.String("clone-cdir") != "" {
			ci.cdir = c.String("clone-cdir")
		}
		if c.String("github-proto") != "" {
			ci.proto = c.String("github-proto")
		}
		ci.quiet = c.Bool("verbose")
		ci.verbose = c.Bool("verbose")

		s, err := NewSeed(ci, target)

		if err != nil {
			println("Unable to retrieve seed")
			println(err.Error())
		}

		err = s.clone()
		fmt.Println(err)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
