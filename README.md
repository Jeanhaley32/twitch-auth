🚧 **Work in Progress** 🚧

# twitch-auth
 Golang Library to obtain [`app access tokens`](https://dev.twitch.tv/docs/authentication/#:~:text=grant%20flow.-,App%20access%20tokens,-APIs%20that%20don%E2%80%99t) using the [client-credentials](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow) Flow for Twitch. 
 ___

### Reason:
  I want to work with the Twitch API on various potential future projects that utilize
  Terminal User Interfaces. To make this easier, I want to create a modular library for
  managing my Oauth2 tokens with Twitch. 

### What This does:
Uses the  flow to retreive, verify, and refresh access tokens.

### What you'll need
You'll first need to [Register your app](https://dev.twitch.tv/docs/authentication/register-app/).
This process should provide you with an `Application ID` and `Secret`.

You can then populate the fields of the TwitchAuth struct, and run the `NewTokenSet()` struct method to populate the token field with
a fresh token. 


### How this works:

Import the twitchauth library

```
import "github.com/jeanhaley32/twitchauth"
```

Create a TwitchAuth struct, and set the Client ID, and secret
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

### Useful Links
[Twitch Authentication Dev Docs](https://dev.twitch.tv/docs/authentication/) - Definition of Token types. 
