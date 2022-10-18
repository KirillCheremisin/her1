package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	jokeResp, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, jokeResp.Joke)
}
