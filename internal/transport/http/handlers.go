package http

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal/repository"
	"main/internal/service"
	"main/internal/transport/http/dto"
	"net/http"
)

var (
	Ok    Status = ResponceStatus{code: http.StatusOK, message: "success"}
	Duble Status = ResponceStatus{code: http.StatusOK, message: "exists"}
	Err   Status = ResponceStatus{code: http.StatusBadRequest, message: "error"}
)

type Handler struct {
	service *service.Service
}

func InitHandlers(db repository.RepositoryInterface) *Handler {
	return &Handler{service: service.InitService(db)}
}
func (h *Handler) SendResponce(w http.ResponseWriter, data any, status Status) {
	fmt.Println(status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.Code())
	if data == nil {
		json.NewEncoder(w).Encode(status.Message())
	} else {
		json.NewEncoder(w).Encode(data)
	}

}
func (h *Handler) CheckConnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodHead {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(Ok.Code())
}
func (h *Handler) GetUserCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.SendResponce(w, nil, Err)
		return
	}

	device := r.PathValue("device")
	count, err := h.service.GetUserCount(r.Context(), device)
	if err != nil {
		h.SendResponce(w, nil, Err)
		return
	}
	h.SendResponce(w, count, Ok)
}
func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.SendResponce(w, nil, Err)
		return
	}

	device := r.PathValue("device")
	index := r.URL.Query().Get("number")

	count, err := h.service.GetUserData(r.Context(), device, index)
	if err != nil {
		h.SendResponce(w, nil, Err)
		return
	}
	h.SendResponce(w, count, Ok)

}
func (h *Handler) AddCardEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.EventDtoIn
	ctx := r.Context()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.SendResponce(w, nil, Err)
	}
	defer r.Body.Close()

	if err = json.Unmarshal(bodyBytes, &req); err != nil {
		err = h.service.SaveErrEvent(ctx, bodyBytes)
		if err != nil {
			h.SendResponce(w, nil, Err)
			return
		}
		h.SendResponce(w, nil, Duble)
		return
	}

	event := req.ToServiceModel()

	if err = h.service.AddCardEvent(ctx, event); err != nil {
		h.SendResponce(w, nil, Err)
		return
	}

	h.SendResponce(w, nil, Ok)
}
