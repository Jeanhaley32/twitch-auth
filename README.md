🚧 **Work in Progress** 🚧

# twitch-auth
 Client Credential Grant Flow for Twitch in Golang
 ___

## Creating a Library for Oauth2 authentication with the Twich API.

### Reason:
  I want to work with the Twitch API on various potential future projects that utilize
  Terminal User Interfaces. To make this easier, I want to create a modular library for
  managing my Oauth2 tokens with Twitch. 

### What This does:
Uses the [client-credentials](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow) flow to retreive, verify, and refresh access tokens.

### How this works:

Import the twitchauth library

```
import "github.com/jeanhaley32/twitchauth"
```

Create a TwitchAuth object, and set the Client ID, and secret
```
	auth := twitchauth.TwitchAuth{
		ClientID: {clientID},
		Secret:   {Secret},
	}
```
Run the NewTokenSet method to obtain a new token. 
```
 auth.NewTokenSet()
```

`.Isexpired()` will return a bool showing if your token is expired.

`.TimeTillExpiration()` Returns a time.duration until you token expires.
