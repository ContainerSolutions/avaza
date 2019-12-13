package main

import (
	"fmt"
	"github.com/moretea/avaza/auth"
	"github.com/moretea/avaza/state"
	"github.com/urfave/cli"
	"log"
)

func loginCommand(c *cli.Context) error {
	bearerToken, err := auth.Register()

	if err != nil {
		fmt.Printf("ERROR LAST")
		return err
	}

	config.BearerToken = bearerToken
	state.SaveConfig(config)
	return nil
}

func accountCommand(c *cli.Context) error {
	authorizer.RefreshIfNecessary()
	accountOk, err := authorizer.Client().Account.AccountGet(nil, authorizer.CreateAuth())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Company:   %s\n", accountOk.Payload.CompanyName)
	fmt.Printf("Subdomain: %s\n", accountOk.Payload.Subdomain)
	return nil
}
