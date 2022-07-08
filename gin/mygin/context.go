package mygin

import "net/http"

type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter
}
