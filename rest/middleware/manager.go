package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for i := len(middlewares) - 1; i >= 0; i-- {
		middleware := middlewares[i]
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WithPrev(handlers http.Handler, middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		n := next

		for _, middleware := range middlewares {
			n = middleware(n)
		}
		return n
	}
}
