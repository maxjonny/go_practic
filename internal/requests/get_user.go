package requests

import (
	"encoding/json"
	"log"
	rep "main/internal/repository"
	"net/http"
)

func GetUserData(w http.ResponseWriter, r *http.Request, db rep.RepositoryInterface) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	device := r.PathValue("device")
	index := r.PathValue("index")

	user, err := db.User.GetUser(device, index)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
