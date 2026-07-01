package http

import (
	"encoding/json"
	"main/internal/repository"
	"net/http"
)

var (
	Ok    Status = ResponceStatus{code: http.StatusOK, message: "success"}
	Duble Status = ResponceStatus{code: http.StatusOK, message: "exists"}
	Err   Status = ResponceStatus{code: http.StatusOK, message: "error"}
)

type Handler struct {
	db repository.RepositoryInterface
}

func InitHandlers(db repository.RepositoryInterface) *Handler {
	return &Handler{db: db}
}

func (h *Handler) SendResponce(w http.ResponseWriter, data any, status Status) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.Code())
	if status.Code() != http.StatusOK {
		json.NewEncoder(w).Encode(status.Message())
	} else {
		json.NewEncoder(w).Encode(data)
	}

}

func (h *Handler) GetUserCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.SendResponce(w, nil, Err)
		return
	}

	device := r.PathValue("device")
	h.db.User.GetUser()
	h.SendResponce(w, nil, Err)

}
