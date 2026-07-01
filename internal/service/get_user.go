package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Service) GetUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	device := r.PathValue("device")
	index := r.PathValue("index")

	user, err := s.rep.User.GetUser(device, index)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
