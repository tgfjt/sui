package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "sui: slack usergroup images"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "usergroup id, g",
			Value: "",
			Usage: "set Slack usergroup id",
		},
		cli.StringFlag{
			Name:  "token, t",
			Value: "",
			Usage: "set Slack API TOKEN",
		},
	}
	app.Action = func(c *cli.Context) error {
		t := c.String("t")
		gID := c.String("g")

		if t == "" {
			log.Fatal("give token!")
		}

		if gID == "" {
			log.Fatal("give a usergroup id!")
		}

		us := GetUsers(t, gID)

		f := func(uID string) {
			GetUser(t, uID).Profile.GetUserImage()
		}

		for _, uID := range us.Users {
			go f(uID)
		}

		time.Sleep(time.Second)

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
