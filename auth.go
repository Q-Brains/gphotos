package gphotos

// AuthorizationScopes represents the authentication scope of PhotosLibraryAPI.
type AuthorizationScopes []authorizationScope

func (scopes *AuthorizationScopes) stringification() []string {
	var results []string
	for _, scope := range *scopes {
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
