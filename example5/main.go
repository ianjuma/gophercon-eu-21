package example5

import (
	"net/http"
)

func server() {
	handlerF := func(w http.ResponseWriter, r1 *http.Request) error {
		return nil
	}
	route := wrapHandlerInMiddleware(handlerF)

	var r2 http.Request
	route(nil, &r2) // BAD: r2 escapes: Indirect Call Flaw
}

type Handler func(w http.ResponseWriter, r3 *http.Request) error

func wrapHandlerInMiddleware(h Handler) Handler {
	f := func(w http.ResponseWriter, r4 *http.Request) error {
		return h(w, r4)
	}
	return f
}
