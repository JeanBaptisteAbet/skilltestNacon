package httphandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"skilltestnacon/database"
)

type Handler struct {
	Context context.Context
	DB      database.DB
}

func (h *Handler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	event, err := h.DB.GetEvent(h.Context, id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(event)
}

func (h *Handler) HandleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	events, err := h.DB.AllActiveEvents(h.Context)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(events)
}
