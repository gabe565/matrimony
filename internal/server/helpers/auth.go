package helpers

import (
	"errors"
	"net/http"
	"strings"
)

const BearerAuthPrefix = "Bearer "

var ErrInvalidBearerAuth = errors.New("invalid bearer auth")

func GetBearerAuthToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if len(authorization) < len(BearerAuthPrefix) || !strings.HasPrefix(authorization, BearerAuthPrefix) {
		return "", ErrInvalidBearerAuth
	}
	return strings.TrimPrefix(authorization, BearerAuthPrefix), nil
}
