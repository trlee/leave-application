package token

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

// securecookie encodes and decodes authenticated and optionally encrypted cookie values.
// Secure cookies can't be forged, because their values are validated using HMAC.
// When encrypted, the content is also inaccessible to malicious eyes.
var CookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// SetCookie encode token into a httponly secure cookie
func SetCookie(token string, w http.ResponseWriter) {
	// Secure cookie
	// Encode token with cookie name as value of the cookie
	if encoded, err := CookieHandler.Encode(CookieName, token); err == nil {
		cookie := &http.Cookie{
			Name:       CookieName,
			Value:      encoded,
			Path:       "/",
			Domain:     "",
			Expires:    ExpiresDate,
			RawExpires: "",
			MaxAge:     MaxAge,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   http.SameSiteStrictMode,
			Raw:        "",
			Unparsed:   []string{},
		}
		http.SetCookie(w, cookie)
	}
}

// Returns decoded cookie (which holds token)
func DecodeCookie(cookie string) (paseto string, err error) {
	if err = CookieHandler.Decode(CookieName, cookie, &paseto); err == nil {
		return paseto, nil
	}
	return "", err
}

// Clear the cookie
func ClearCookie(w http.ResponseWriter) {
	log.Printf("%s %s %s \n", "the cookie", CookieName, "have been deleted successfully")
	cookie := &http.Cookie{
		Name:   CookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
