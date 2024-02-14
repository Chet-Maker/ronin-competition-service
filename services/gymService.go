package services

import (
	"encoding/json"
	"log"
	"net/http"
	"ronin/models"
	"ronin/repositories"

	"github.com/gorilla/mux"
	// "github.com/jmoiron/sqlx"
)

var gymRepo *repositories.GymRepository

type GymService struct {
	gymRepo *repositories.GymRepository
}

func NewGymService(gymRepo *repositories.GymRepository) *GymService {
	return &GymService{
		gymRepo: gymRepo,
	}
}

func (g *GymService) GetAllGyms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gyms = models.GetGyms()
	gyms, err := g.gymRepo.GetAllGyms()
	if err == nil {
		json.NewEncoder(w).Encode(&gyms)
	} else {
		log.Println(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

}

func (g *GymService) CreateGym(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gym = models.GetGym()
	_ = json.NewDecoder(r.Body).Decode(&gym)
	gym, err := g.gymRepo.CreateGym(gym)
	if err == nil {
		json.NewEncoder(w).Encode(&gym)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (g *GymService) GetGym(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gyms = models.GetGym()
	vars := mux.Vars(r)
	id := vars["gym_id"]
	gyms, err := g.gymRepo.GetGymById(id)
	if err == nil {
		json.NewEncoder(w).Encode(&gyms)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}
