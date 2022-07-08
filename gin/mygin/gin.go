package mygin

import "net/http"

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

// RoutesInfo defines a RouteInfo slice.
type RoutesInfo []RouteInfo

type Engine struct {
	trees methodTrees
}

func New() *Engine {
	engine := &Engine{}

	return engine
}

func (engine *Engine) ServerHTTP(w http.ResponseWriter, req *http.Request) {}

func (engine *Engine) handleHTTPRequest(c *Context) {

}

func (engine *Engine) HandleContext(c *Context) {

	engine.handleHTTPRequest(c)

}

func (engine *Engine) Handler() http.Handler {

}
