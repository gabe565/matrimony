package helpers

import (
	"fmt"
	"net/http"
)

func RequestProto(r *http.Request) string {
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	}
	return "http"
}

func PublicUrl(r *http.Request) string {
	return fmt.Sprintf("%s://%s", RequestProto(r), r.Host)
}
