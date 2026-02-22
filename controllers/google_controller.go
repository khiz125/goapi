package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/khiz125/goapi/config"
	"github.com/khiz125/goapi/services/auth"
)

type GoogleAuthController struct {
	oauthConfig config.GoogleAuthConfig
	authService *auth.AuthService
}

func NewGoogleAuthController(cfg config.GoogleAuthConfig, authService *auth.AuthService) *GoogleAuthController {
	return &GoogleAuthController{
		oauthConfig: cfg,
		authService: authService,
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

	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, "state cookier not found", http.StatusBadRequest)
		return
	}

	if cookie.Value != state {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "oauth_state",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	user, err := c.authService.HandleGoogleCallback(r.Context(), code)
	if err != nil {
		http.Error(w, "auth failed", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello " + user.Name))
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

	return config.GoogleAuthURL + "?" + v.Encode()
}
