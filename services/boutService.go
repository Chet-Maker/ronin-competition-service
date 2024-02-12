package services

import (
	"encoding/json"
	"net/http"
	"ronin/models"
	"ronin/repositories"

	"github.com/gorilla/mux"
)

var boutRepo *repositories.BoutRepository

type BoutService struct {
	boutRepo *repositories.BoutRepository
}

func NewBoutService(boutRepo *repositories.BoutRepository) *BoutService {
	return &BoutService{
		boutRepo: boutRepo,
	}
}

func (b *BoutService) GetAllBouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bouts = models.GetBouts()
	bouts, err := boutRepo.GetAllBouts()
	if err == nil {
		json.NewEncoder(w).Encode(&bouts)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (b *BoutService) GetBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bouts = models.GetBout()
	vars := mux.Vars(r)
	id := vars["bout_id"]
	bouts, err := b.boutRepo.GetBoutById(id)
	if err == nil {
		json.NewEncoder(w).Encode(&bouts)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (b *BoutService) CreateBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bout = models.GetBout()
	_ = json.NewDecoder(r.Body).Decode(&bout)
	if bout.ChallengerId != bout.AcceptorId {
		boutId, err := b.boutRepo.CreateBout(bout)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var outboundBout models.OutboundBout
		outboundBout, err = b.boutRepo.GetOutboundBoutByBoutId(boutId)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		json.NewEncoder(w).Encode(&outboundBout)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := map[string]string{
			"error": "ChallengerId and AcceptorId must be different. You cannot create a bout against yourself.",
		}
		json.NewEncoder(w).
			Encode(errorMessage)
	}
}

func (b *BoutService) UpdateBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bout = models.GetBout()
	_ = json.NewDecoder(r.Body).Decode(&bout)
	vars := mux.Vars(r)
	id := vars["bout_id"]
	err := b.boutRepo.UpdateBout(id, bout)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&bout)
}

func (b *BoutService) DeleteBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["bout_id"]
	err := b.boutRepo.DeleteBout(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&id)
}

func (b *BoutService) AcceptBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["bout_id"]
	err := b.boutRepo.AcceptBout(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&id)
}

func (b *BoutService) DeclineBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["bout_id"]
	err := b.boutRepo.DeclineBout(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&id)
}

func (b *BoutService) CompleteBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	boutId := vars["bout_id"]
	refereeId := vars["referee_id"]
	err := b.boutRepo.CompleteBout(boutId, refereeId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&boutId)
}

func (b *BoutService) GetPendingBouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bouts []models.OutboundBout
	vars := mux.Vars(r)
	id := vars["athlete_id"]
	bouts, err := b.boutRepo.GetPendingBoutsByAthleteId(id)
	if err == nil {
		json.NewEncoder(w).Encode(&bouts)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (b *BoutService) GetIncompleteBouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bouts []models.OutboundBout
	vars := mux.Vars(r)
	id := vars["athlete_id"]
	bouts, err := b.boutRepo.GetIncompleteBoutsByAthleteId(id)
	if err == nil {
		json.NewEncoder(w).Encode(&bouts)
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (b *BoutService) CancelBout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	boutId := vars["bout_id"]
	athleteId := vars["challenger_id"]
	err := b.boutRepo.CancelBout(boutId, athleteId)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&boutId)
}
