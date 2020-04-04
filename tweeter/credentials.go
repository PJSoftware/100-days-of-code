package tweeter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const tokensFile string = "//PERSONALCLOUD/Public/Tokens/API/twitter.json"

type token struct {
	Name  string        `json:"name"`
	Creds []credentials `json:"credentials"`
}

type credentials struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

func newCredentials(appName string) *credentials {
	cred := new(credentials)
	tf := chooseTokens()
	fmt.Printf("Reading '%s' credentials from '%s'\n", appName, tf)
	jsonFile, err := os.Open(tf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(byteValue), &jsonData)

	cred.parseToken(appName, jsonData["tokens"])
	return cred
}

func chooseTokens() string {
	files := []string{"../Tokens/API/twitter.json"}
	for _, tf := range files {
		if fileExists(tf) {
			return tf
		}
	}
	return tokensFile
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (cred *credentials) parseToken(appName, data interface{}) {
	for _, token := range data.([]interface{}) {
		tm := token.(map[string]interface{})
		if tm["name"] == appName {
			ci := tm["credentials"]
			cs := ci.([]interface{})
			csi := cs[0]
			cm := csi.(map[string]interface{})
			cred.consumerKey = cm["CONSUMER_KEY"].(string)
			cred.consumerSecret = cm["CONSUMER_SECRET"].(string)
			cred.accessToken = cm["ACCESS_TOKEN"].(string)
			cred.accessSecret = cm["ACCESS_TOKEN_SECRET"].(string)
		}
	}
}
