package mygin

import "net/http"

type ResponseWriter interface {
	http.ResponseWriter

	WriteString(string) (int, error)
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}
