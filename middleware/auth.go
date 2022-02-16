package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
}

func Auth(w http.ResponseWriter, req *http.Request) string {
	//func Auth(req *http.Request) string {
	client := &http.Client{}
	url := "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"
	//fmt.Println("APP_ID", os.Getenv("APP_ID"))
	//fmt.Println("Secret_key", os.Getenv("SECRET_KEY"))
	x := fmt.Sprintln("grant_type=client_credentials&client_id=", os.Getenv("APP_ID"), "&client_secret=", os.Getenv("SECRET_KEY"), "&scope=https%3A%2F%2Fapi.botframework.com%2F.default")
	//fmt.Printf("Value of x: %v", x)
	payload := strings.NewReader(x)
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("NewRequest failed")
	}
	req.Header.Add("Host", "login.microsoftonline.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf("Processing Request Failed")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var AccessToken Token
	err = json.Unmarshal(body, &AccessToken)
	if err != nil {
		return fmt.Sprintf("Unmarshalling Failed")
	}
	fmt.Println(string(body))

	return AccessToken.AccessToken
}
