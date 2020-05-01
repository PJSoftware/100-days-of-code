package tweeter_test

import (
	"testing"

	"github.com/PJSoftware/TweetCommit/tweeter"
)

func TestGetCredentials(t *testing.T) {
	_, err := tweeter.GetCredentials("API/twitter.json", "PJ_AutoTweeter")
	if err != nil {
		t.Errorf("Error reading required credentials: %s", err)
	}

	cred, _ := tweeter.GetCredentials("TEST/twitter-api.json", "Twitter_App_One")
	if cred.ConsKey != "abcde12345" {
		t.Errorf("ConsKey not allocated correctly")
	}
	if cred.ConsSec != "bcdef23456" {
		t.Errorf("ConsSec not allocated correctly")
	}
	if cred.AccTok != "cdefg34567" {
		t.Errorf("AccTok not allocated correctly")
	}
	if cred.AccTokSec != "defgh45678" {
		t.Errorf("AccTokSec not allocated correctly")
	}

}
