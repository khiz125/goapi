package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/khiz125/goapi/apperrors"
	"github.com/khiz125/goapi/common"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = `589645437572-r19negktr11k6gm1qsjo54cg4eqmklab.apps.googleusercontent.com`
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if strings.HasPrefix(req.URL.Path, "/auth/") {
			next.ServeHTTP(w, req)
      return
		}

		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Fields(authorization)
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid request Header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]

		if !strings.EqualFold(bearer, "Bearer") || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid request header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		ctx := req.Context()
		tokenValidator, err := idtoken.NewValidator(ctx)
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(ctx, idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]

		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
