package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"os/user"
)

func main() {
	app := cli.NewApp()
	app.Name = "seeder"
	app.Usage = "copy source from remote repository"
	app.Version = "0.0.1"

	usr, _ := user.Current()
	defaultKey := usr.HomeDir + "/.ssh/id_rsa"

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
			Name:  "branch, b",
			Usage: "Git branch",
		},
		cli.StringFlag{
			Name:     "clone-cdir, c",
			Usage:    "temp directory where repo will be cloned; \"memory\" to use system memory",
			FilePath: "/tmp",
			Value:    "/tmp",
		},
		cli.StringFlag{
			Name:  "ssh-key, k",
			Usage: "private ssh key to use (default: \"" + defaultKey + "\")",
		},
		cli.StringFlag{
			Name:  "github-proto, p",
			Usage: "Github proto; auto, https, or ssh",
			Value: "auto",
		},
		cli.BoolFlag{
			Name:  "preserve-git, g",
			Usage: "preserve .git directory",
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
		ci := ConfigInput{}
		ci.target = c.Args().Get(0)
		ci.src = c.Args().Get(1)
		ci.dst = c.Args().Get(2)
		ci.branch = c.String("branch")
		ci.cdir = c.String("clone-cdir")
		ci.key = c.String("ssh-key")
		ci.proto = c.String("github-proto")
		ci.git = c.Bool("preserve-git")
		ci.quiet = c.Bool("quiet")
		ci.verbose = c.Bool("verbose")

		s, err := NewSeed(ci)
		if err != nil {
			println("Unable to retrieve seed")
			println(err.Error())
		}

		err = s.clone()
		if err != nil {
			println("Unable to clone project")
			println(err.Error())
		}

		err = s.process()
		if err != nil {
			println("Unable to copy files")
			println(err.Error())
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
