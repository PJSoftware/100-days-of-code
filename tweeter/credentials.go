package tweeter

import "github.com/pjsoftware/gotokens"

// Credentials store out twitter auth details
type Credentials struct {
	ConsKey   string
	ConsSec   string
	AccTok    string
	AccTokSec string
}

// GetCredentials returns the Twitter login credentials
func GetCredentials(jsonFile, appName string) (*Credentials, error) {
	op := "tweeter.GetCredentials"
	cred := new(Credentials)

	gotokens.SetSearchPath([]string{"C:", "Z:"})
	tks, err := gotokens.ImportTokens(jsonFile)
	if err != nil {
		return nil, &gotokens.Error{Op: op, Err: err}
	}
	tk, err := tks.Select(appName)
	if err != nil {
		return nil, &gotokens.Error{Op: op, Err: err}
	}

	for _, tok := range []string{
		"CONSUMER_KEY", "CONSUMER_SECRET",
		"ACCESS_TOKEN", "ACCESS_TOKEN_SECRET",
	} {
		v, err := tk.Credential(tok)
		if err != nil {
			return nil, &gotokens.Error{Op: op, Err: err}
		}
		switch tok {
		case "CONSUMER_KEY":
			cred.ConsKey = v
		case "CONSUMER_SECRET":
			cred.ConsSec = v
		case "ACCESS_TOKEN":
			cred.AccTok = v
		case "ACCESS_TOKEN_SECRET":
			cred.AccTokSec = v
		}
	}

	return cred, nil
}
