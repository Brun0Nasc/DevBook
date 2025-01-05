package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configure uses the environment variables to create the SecureCookie
func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save records authentication information
func Save(w http.ResponseWriter, id, token string) (err error) {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return
}

// Read returns values stored in cookies
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
