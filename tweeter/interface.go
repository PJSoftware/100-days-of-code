package tweeter

import (
	"fmt"

	"github.com/PJSoftware/TweetCommit/logdata"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const url string = "https://github.com/PJSoftware/100-days-of-code/blob/master/log.md"

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
func TweetHDC(ld *logdata.LogData) error {
	creds := newCredentials("PJ_AutoTweeter")
	client, err := getClient(creds)
	if err != nil {
		return fmt.Errorf("Could not initialise twitter client: %w", err)
	}
	_ = client

	tweet := fmt.Sprintf("#100DaysOfCode #R%dD%d Day %d\n", ld.Round, ld.Day, ld.Day)
	tweet += ld.Topic + ": " + ld.Desc + "\n"
	tweet += url

	fmt.Printf("Tweeting the following:\n=====\n%s\n=====\n", tweet)

	_, _, err = client.Statuses.Update(tweet, nil)
	return err
}
