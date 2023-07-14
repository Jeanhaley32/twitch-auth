// Library for authenticating with Twitch using OAuth2 and the client credentials grant flow
// https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow
package twitchauth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Constants for the Twitch API
const (
	// Twitch API URL
	twitchAuthTokenURL = "https://id.twitch.tv/oauth2/token"
)

// TwitchAuth is the struct for the Twitch API
type TwitchAuth struct {
	ClientID string
	Secret   string
	Token    token
}

// token is the response from the Twitch API
type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

type TwitchAuthInterface interface {
	NewTokenSet()
}

// Obtains a new Token set from the Twitch API
// Token set includes access toke, and refresh token
func (self TwitchAuth) NewTokenSet() error {
	var t token
	// Client credentials grant flow
	// https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow
	data := url.Values{}
	data.Set("client_id", self.ClientID)
	data.Set("client_secret", self.Secret)
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", twitchAuthTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// // (test) Print the whole request to the console for debugging
	// dump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	return err
	// }

	// fmt.Printf("\nRequest: %s\n", dump)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// // (test) Print Server Response
	// dump, err = httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("\nResponse: %s\n", dump)

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return err
	}

	return nil
}
