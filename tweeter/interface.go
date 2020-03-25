package tweeter

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const url string = "https://github.com/PJSoftware/100-days-of-code/blob/master/log.md"
const testOnly bool = true

func getClient(creds *credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.consumerKey, creds.consumerSecret)
	token := oauth1.NewToken(creds.accessToken, creds.accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// TweetHDC adds the #100DaysofCode details around our message, and tweets it
func TweetHDC(round int, day int, msg string) error {
	creds := newCredentials("PJ_AutoTweeter")
	client, err := getClient(creds)
	if err != nil {
		return fmt.Errorf("Could not initialise twitter client: %w", err)
	}
	_ = client

	tweet := fmt.Sprintf("#100DaysOfCode #R%dD%d Day %d\n", round, day, day)
	tweet += msg + "\n"
	tweet += url

	if testOnly {
		fmt.Printf("\n=====\nTEST ONLY; nothing will be sent to twitter!\nTweet:\n%s\n=====\n\n", tweet)
		return nil
	}

	_, _, err = client.Statuses.Update(tweet, nil)
	return err
}
