// Library for authenticating with Twitch using OAuth2 and the client credentials grant flow
// https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow
package twitchauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Constants for the Twitch API
const (
	// Twitch API URL
	twitchURLPrefix = "https://api.twitch.tv/"
	// Twitch API URL sub-domain for getting a token
	tokenURl = "oauth2/token"
)

// TwitchAuth is the struct for the Twitch API
type TwitchAuth struct {
	ClientID string
	Secret   string
	Token    token
}

// token is the response from the Twitch API
type token struct {
	AccessToken   string `json:"access_token"`
	TokenType     string `json:"token_type"`
	ExpiresIn     int64  `json:"expires_in"`
	Scope         string `json:"scope"`
	Refresh_token string `json:"refresh_token"`
}

type TwitchAuthInterface interface {
	NewTokenSet()
}

func (self TwitchAuth) NewTokenSet() error {
	// Client credentials grant flow
	// https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow
	data := url.Values{}
	data.Set("client_id", self.ClientID)
	data.Set("client_secret", self.Secret)
	data.Set("grant_type", "client_credentials")
	fmt.Println("URL: " + twitchURLPrefix + tokenURl + data.Encode())
	req, err := http.NewRequest("POST", twitchURLPrefix+tokenURl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var t token

	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return err
	}
	fmt.Printf("access token: %v\n", t.AccessToken)
	self.Token = t
	return nil
}
