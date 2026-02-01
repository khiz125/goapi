package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/khiz125/goapi/config"
)

type GoogleAuthController struct {
	oauthConfig config.GoogleAuthConfig
}

func NewGoogleAuthController(cfg config.GoogleAuthConfig) *GoogleAuthController {
	return &GoogleAuthController{
		oauthConfig: cfg,
	}
}

func (c *GoogleAuthController) Login(w http.ResponseWriter, r *http.Request) {
	state := generateState()

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	authURL := c.buildGoogleAuthURL(state)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func (c *GoogleAuthController) CallBack(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query.Get("code")
	state := query.Get("state")

	if code == "" || state == "" {
		http.Error(w, "missing code or state", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("auth_state")
	if err != nil {
		http.Error(w, "state cookier not found", http.StatusBadRequest)
		return
	}

	if cookie.Value != state {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
  w.Write([]byte("state validation OK"))  // TODO: call service
}

func generateState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawStdEncoding.EncodeToString(b)
}

func (c *GoogleAuthController) buildGoogleAuthURL(state string) string {
	v := url.Values{}
	v.Set("client_id", c.oauthConfig.ClientID)
	v.Set("redirect_uri", c.oauthConfig.RedirectURI)
	v.Set("response_type", "code")
	v.Set("scope", "openid profile email")
	v.Set("state", state)

	return c.oauthConfig.AuthURL + "?" + v.Encode()
}
