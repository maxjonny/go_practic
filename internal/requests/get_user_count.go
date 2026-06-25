package requests

import (
	"encoding/json"
	"log"
	db "main/internal/database"
	m "main/internal/models"
	"net/http"
)

func GetUserCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var users []m.UserCard
	var err error
	device := r.PathValue("device")

	nodeIds, err := db.Storage.Rep.Device.GetActiveNode(device)
	if err != nil {
		log.Println(err)
	}

	if len(nodeIds) > 0 {
		users, err = db.Storage.Rep.User.GetUserByNodes(nodeIds)
		if err != nil {
			log.Println(err)
		}
	}

	if len(users) > 0 {
		db.Storage.Rep.User.DropCache(device)
		db.Storage.Rep.User.CreateCache(device, users)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(len(users))
}
