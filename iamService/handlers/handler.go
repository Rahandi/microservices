package handlers

import "net/http"

type Handler interface {
	Register(httpServer *http.ServeMux)
}
