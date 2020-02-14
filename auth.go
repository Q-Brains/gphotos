package gphotos

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (auth authorizations) OAuth2InteractiveFlow(clientID string, clientSecret string, scopes AuthorizationScopes, state string, options ...oauth2.AuthCodeOption) (*http.Client, error) {
	conf := Auth.OAuth2Config(clientID, clientSecret, scopes)
	authURL := Auth.OAuth2CreateURL(conf, state, options...)

	fmt.Println("Access the following URL and paste the Authorization Code.")
	fmt.Println("> URL: " + authURL)
	fmt.Print("> Authorization Code: ")

	var s string
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	return Auth.OAuth2CreateClient(conf, s)
}

func (auth authorizations) OAuth2Config(clientID string, clientSecret string, scopes AuthorizationScopes) oauth2.Config {
	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       scopes.stringification(),
	}
}

func (auth authorizations) OAuth2CreateURL(conf oauth2.Config, state string, options ...oauth2.AuthCodeOption) string {
	return conf.AuthCodeURL(state, options...)
}

func (auth authorizations) OAuth2CreateClient(conf oauth2.Config, authCode string) (*http.Client, error) {
	token, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext, token), nil
}

// AuthorizationScopes represents the authentication scope of PhotosLibraryAPI.
type AuthorizationScopes []authorizationScope

func (scopes AuthorizationScopes) stringification() []string {
	var results []string
	for _, scope := range scopes {
		results = append(results, string(scope))
	}
	return results
}

type authorizationScope string

// Here's the authorization scopes.
// Source: https://developers.google.com/photos/library/guides/authentication-authorization#OAuth2Authorizing
const (
	// Read access only.
	// List items from the library and all albums, access all media items and list albums owned by the user, including those which have been shared with them.
	// For albums shared by the user, share properties are only returned if the .sharing scope has also been granted.
	// The ShareInfo property for albums and the contributorInfo for mediaItems is only available if the .sharing scope has also been granted.
	// For more information, see Share media.
	Readonly authorizationScope = "https://www.googleapis.com/auth/photoslibrary.readonly"

	// Write access only.
	// Acess to upload bytes, create media items, create albums, and add enrichments. Only allows new media to be created in the user's library and in albums created by the app.
	Appendonly = "https://www.googleapis.com/auth/photoslibrary.appendonly"

	// Read access to media items and albums created by the developer. For more information, see Access media items and List library contents, albums, and media items.
	// Intended to be requested together with the Appendonly scope.
	Appcreateddata = "https://www.googleapis.com/auth/photoslibrary.readonly.appcreateddata"

	// Access to both the Appendonly and Readonly scopes. Doesn't include Sharing.
	ReadAndAppend = "https://www.googleapis.com/auth/photoslibrary"

	// Access to sharing calls.
	// Access to create an album, share it, upload media items to it, and join a shared album.
	Sharing = "https://www.googleapis.com/auth/photoslibrary.sharing"
)
