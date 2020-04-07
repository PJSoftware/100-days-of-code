package tweeter

import (
	"fmt"

	"github.com/PJSoftware/TweetCommit/logdata"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/pjsoftware/gotokens"
)

const twitterApp string = "PJ_AutoTweeter"

// Tweeter is our struct of critical data
type Tweeter struct {
	client *twitter.Client
	Err    error
}

// NewTweeter sets up our twitter interface
func NewTweeter() *Tweeter {
	tw := new(Tweeter)

	gotokens.SetSearchPath([]string{"C:", "Z:"})
	tks, err := gotokens.ImportTokens("API/twitter.json")
	if err != nil {
		tw.Err = err
		return nil
	}
	tk, err := tks.Select("PJ_AutoTweeter")
	if err != nil {
		tw.Err = err
		return nil
	}

	ck, e1 := tk.Credential("CONSUMER_KEY")
	cs, e2 := tk.Credential("CONSUMER_SECRET")
	at, e3 := tk.Credential("ACCESS_TOKEN")
	as, e4 := tk.Credential("ACCESS_TOKEN_SECRET")
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		tw.Err = &gotokens.Error{
			Code:    gotokens.EBADCREDENTIAL,
			Message: "Credential unrecognised",
		}
		return nil
	}

	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	tw.getClient(config, token)

	return tw
}

func (tw *Tweeter) getClient(config *oauth1.Config, token *oauth1.Token) {
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
