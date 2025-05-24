package main

import (
	"context"
	"fmt"
	"github.com/andygrunwald/go-gerrit"
)

func main() {
	instance := "https://gerrit.teimas.com/"
	ctx := context.Background()
	client, _ := gerrit.NewClient(ctx, instance, nil)
	client.Authentication.SetDigestAuth("victorps", "")
	self, _, _ := client.Accounts.GetAccount(ctx, "victorps")

	fmt.Printf("Username: %s", self.Name)
}