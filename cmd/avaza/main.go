package main

import (
	authorizer_ "github.com/ContainerSolutions/avaza/authorizer"
	"github.com/ContainerSolutions/avaza/state"
	"github.com/urfave/cli"
	"log"
	"os"
)

var config *state.Config
var authorizer *authorizer_.Authorizer

func init() {
	config = state.LoadConfig()
	authorizer = authorizer_.NewAuthorizer(config)
}

func main() {

	app := cli.NewApp()
	app.Name = "Avaza"
	app.Usage = "track projects, tasks & time"

	app.Commands = []cli.Command{
		{
			Name:   "login",
			Usage:  "(re)login in Avaza",
			Action: loginCommand,
		},
		{
			Name:   "account",
			Usage:  "Get account info",
			Action: accountCommand,
		},
		{
			Name:    "project",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  projectListCommand,
				},
			},
		},
	}

	_ = config

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
