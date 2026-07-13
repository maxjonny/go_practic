package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/internal/repository"
	"main/internal/service"
	"main/internal/transport/dto"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *Handler) SendResponce(w http.ResponseWriter, data any, status Status, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		data = Err.Message()
	} else {
		if data == nil {
			data = status.Message()
		}
	}

	w.WriteHeader(status.Code())
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) CheckConnect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(Ok.Code())
}

func (h *Handler) GetUserCount(w http.ResponseWriter, r *http.Request) {

	device := r.PathValue("device")
	count, err := h.service.GetUserCount(r.Context(), device)
	if err != nil {
		h.SendResponce(w, nil, Err, fmt.Errorf("GetUserCount: %w", err))
		return
	}

	w.Write([]byte(strconv.Itoa(count)))
}

func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {

	device := r.PathValue("device")
	index := r.URL.Query().Get("number")

	user, err := h.service.GetUserData(r.Context(), device, index)
	if err != nil {
		h.SendResponce(w, nil, Err, fmt.Errorf("GetUserData: %w", err))
		return
	}

	var userDto dto.UserDto
	res := make([]dto.UserDto, 0, 1)

	userDto.FromServiceModel(user)
	res = append(res, userDto)

	h.SendResponce(w, res, Ok, nil)

}

func (h *Handler) AddCardEvent(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	var req dto.EventDtoIn

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.SendResponce(w, nil, Err, fmt.Errorf("ReadAll: %w", err))
	}
	defer r.Body.Close()

	err = json.Unmarshal(bodyBytes, &req)
	if err != nil || !req.IsValid() {
		err = h.service.SaveErrEvent(ctx, bodyBytes)
		if err != nil {
			h.SendResponce(w, nil, Err, fmt.Errorf("SaveErrEvent: %w", err))
			return
		}
		h.SendResponce(w, nil, Duble, nil)
		return
	}

	event := req.ToServiceModel()

	isAded, err := h.service.AddCardEvent(ctx, event)
	if err != nil {
		h.SendResponce(w, nil, Err, fmt.Errorf("AddCardEvent: %w", err))
		return
	}

	if !isAded {
		h.SendResponce(w, nil, Duble, nil)
		return
	}

	h.SendResponce(w, nil, Ok, nil)
}

func (h *Handler) UploadUser(w http.ResponseWriter, r *http.Request) {

	h.SendResponce(w, nil, Duble, nil)

	// var req dto.UserDto
	// ctx := context.Background()

	// bodyBytes, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	h.SendResponce(w, nil, Err, fmt.Errorf("ReadAll: %w", err))
	// }
	// defer r.Body.Close()

	// err = json.Unmarshal(bodyBytes, &req)
	// if err != nil || !req.IsValid() {
	// 	h.SendResponce(w, nil, Duble, nil)
	// 	return
	// }

	// user := req.ToServiceModel()

	// isAded, err := h.service.UpsertUser(ctx, user)
	// if err != nil {
	// 	h.SendResponce(w, nil, Err, fmt.Errorf("UpsertUser: %w", err))
	// 	return
	// }

	// if !isAded {
	// 	h.SendResponce(w, nil, Duble, nil)
	// 	return
	// }

	//h.SendResponce(w, nil, Ok, nil)
}

func (h *Handler) UpdateConnect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		device := chi.URLParam(r, "device")
		if err := h.service.UpdateBoxStatus(context.Background(), device); err != nil {
			h.SendResponce(w, "Unknown device: "+device, Err, fmt.Errorf("UpdateBoxStatus: %w", err))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func (h *Handler) Test(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		log.Println(r.URL.Path)

// 		next.ServeHTTP(w, r)
// 	})
// }
