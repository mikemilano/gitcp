package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	// TODO: Add option for ssh key
	// TODO: Add option for github https vs ssh

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "target",
			Usage: "Github user/project or any repo path",
		},
	}

	app.Action = func(c *cli.Context) error {
		target := c.Args().Get(0)

		s, err := NewSeed(target)

		if err != nil {
			println("Unable to retrieve seed")
		}

		s.clone()


		//fmt.Println("Hola", target)

		//clone()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
