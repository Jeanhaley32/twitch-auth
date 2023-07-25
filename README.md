🚧 **Work in Progress** 🚧

# twitch-auth
 Golang Library to obtain an [`app access token`](https://dev.twitch.tv/docs/authentication/#:~:text=grant%20flow.-,App%20access%20tokens,-APIs%20that%20don%E2%80%99t) using the [`client-credentials`](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth#oauth-client-credentials-flow) work Flow for Twitch. 
 ___

### Useful Links
- [Twitch Authentication Dev Docs](https://dev.twitch.tv/docs/authentication/) - Definition of Token types.
- [Oauth 2.0 White paper](https://datatracker.ietf.org/doc/html/rfc6749) - Informative doc on what & how Oauth2.0 actually works.
- [Twitch Authentication flows diagram](https://dev.twitch.tv/docs/authentication/#:~:text=grant%20flow.-,Authentication%20flows,-The%20following%20table) - useful at a glance difference between Token types.
- [Twitch Client Credentials Workflow](https://dev.twitch.tv/docs/authentication/getting-tokens-oauth/#client-credentials-grant-flow) - Twithc page for the Client Credentials Grant flow. 



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

