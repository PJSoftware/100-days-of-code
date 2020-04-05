package tweeter

import (
	"fmt"

	"github.com/PJSoftware/TweetCommit/logdata"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const twitterApp string = "PJ_AutoTweeter"

// Tweeter is our struct of critical data
type Tweeter struct {
	client *twitter.Client
	cred   *credentials
	Err    error
}

// NewTweeter sets up our twitter interface
func NewTweeter() *Tweeter {
	tw := new(Tweeter)
	tw.cred = newCredentials(twitterApp)
	tw.getClient()
	return tw
}

func (tw *Tweeter) getClient() {
	config := oauth1.NewConfig(tw.cred.consumerKey, tw.cred.consumerSecret)
	token := oauth1.NewToken(tw.cred.accessToken, tw.cred.accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		tw.Err = err
		return
	}
	tw.client = client
}

const url string = "https://github.com/PJSoftware/100-days-of-code/blob/master/log.md"

// TweetHDC adds the #100DaysofCode details around our message, and tweets it
func (tw *Tweeter) TweetHDC(ld *logdata.LogData) error {
	tweet := fmt.Sprintf("#100DaysOfCode #R%dD%d Day %d\n", ld.Round, ld.Day, ld.Day)
	tweet += ld.Topic + ": " + ld.Desc + "\n"
	tweet += url

	fmt.Printf("Tweeting the following:\n=====\n%s\n=====\n", tweet)

	_, _, err := tw.client.Statuses.Update(tweet, nil)
	return err
}
