package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	twClient = new(Twitter)
)

// Twitter holds the twitter client struct
type Twitter struct {
	Client *twitter.Client
}

func setupTwitter() {
	if os.Getenv("SKIP_TWITTER") == "true" {
		return
	}
	// Login and stuff
	twconfig := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))
	httpClient := twconfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		fmt.Printf("[Error] %s\n", err)
	}
	twClient.Client = client
}
