package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ronin/repositories"

	"github.com/gorilla/mux"
)

var feedRepo *repositories.FeedRepository

type FeedService struct {
	feedRepo *repositories.FeedRepository
}

func NewFeedService(feedRepo *repositories.FeedRepository) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
	}
}
func (f *FeedService) GetFeedByAthleteId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["athlete_id"]

	feed, err := f.feedRepo.GetFeedByAthleteId(id)

	switch err {
	case sql.ErrNoRows:
		log.Println("no rows returns.")
		http.Error(w, err.Error(), 204)
	case nil:
		json.NewEncoder(w).Encode(&feed)
	default:
		http.Error(w, err.Error(), 400)
		return
	}
}
