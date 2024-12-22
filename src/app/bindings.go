package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Bind(r *mux.Router, h *Handler) {
	r.Path("/send_message").Methods(http.MethodPost).
		HandlerFunc(h.HandlePostMessage)
}
