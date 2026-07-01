package service

import (
	"encoding/json"
	"log"
	m "main/internal/models"
	"net/http"
)

func (s *Service) GetUserCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var users []m.UserCard
	var err error
	device := r.PathValue("device")

	nodeIds, err := s.rep.Device.GetActiveNode(device)
	if err != nil {
		log.Println(err)
	}

	if len(nodeIds) > 0 {
		users, err = s.rep.User.GetUserByNodes(nodeIds)
		if err != nil {
			log.Println(err)
		}
	}

	if len(users) > 0 {
		s.rep.User.DropCache(device)
		s.rep.User.CreateCache(device, users)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(len(users))
}
