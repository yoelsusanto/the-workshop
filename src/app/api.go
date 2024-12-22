package app

import (
	"log/slog"
	"net/http"

	"github.com/yoelsusanto/the-workshop/src/awssdk"
)

type Handler struct {
	MainQueue *awssdk.SQSClient
}

func (h *Handler) HandlePostMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.MainQueue.SendMessage(ctx, "Hello world!"); err != nil {
		slog.Error("Failed to send message to main queue.", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
