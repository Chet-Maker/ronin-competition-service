package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"

	"ronin/services"
)

const base_url = "/api/v1"

func CreateRouter(db *sqlx.DB,
	athleteService *services.AthleteService,
	feedService *services.FeedService,
	styleService *services.StyleService,
	boutService *services.BoutService,
	outcomeService *services.OutcomeService,
	athleteScoreService *services.AthleteScoreService,
	gymService *services.GymService) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(base_url+"/athletes", athleteService.GetAllAthletes).Methods("GET")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteService.GetAthlete).Methods("GET")
	router.HandleFunc(base_url+"/athlete", athleteService.CreateAthlete).Methods("POST")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteService.UpdateAthlete).Methods("PUT")
	router.HandleFunc(base_url+"/athlete/{athlete_id}", athleteService.DeleteAthlete).Methods("DELETE")
	router.HandleFunc(base_url+"/athlete/all/usernames", athleteService.GetAllAthleteUsernames).Methods("GET")
	router.HandleFunc(base_url+"/athlete/{athlete_id}/record", athleteService.GetAthleteRecord).Methods("GET")
	router.HandleFunc(base_url+"/athlete/authorize", athleteService.IsAuthorizedUser).Methods("POST")
	router.HandleFunc(base_url+"/athletes/follow", athleteService.FollowAthlete).Methods("POST")
	router.HandleFunc(base_url+"/athletes/{followerId}/{followedId}/unfollow", athleteService.UnfollowAthlete).Methods("DELETE")
	router.HandleFunc(base_url+"/athletes/following/{athlete_id}", athleteService.GetAthletesFollowed).Methods("GET")

	router.HandleFunc(base_url+"/bouts", boutService.GetAllBouts).Methods("GET")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutService.GetBout).Methods("GET")
	router.HandleFunc(base_url+"/bout", boutService.CreateBout).Methods("POST")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutService.UpdateBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}", boutService.DeleteBout).Methods("DELETE")
	router.HandleFunc(base_url+"/bout/{bout_id}/accept", boutService.AcceptBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}/decline", boutService.DeclineBout).Methods("PUT")
	router.HandleFunc(base_url+"/bout/{bout_id}/complete/{referee_id}", boutService.CompleteBout).Methods("PUT")
	router.HandleFunc(base_url+"/bouts/pending/{athlete_id}", boutService.GetPendingBouts).Methods("GET")
	router.HandleFunc(base_url+"/bouts/incomplete/{athlete_id}", boutService.GetIncompleteBouts).Methods("GET")
	router.HandleFunc(base_url+"/bout/cancel/{bout_id}/{challenger_id}", boutService.CancelBout).Methods("PUT")

	router.HandleFunc(base_url+"/gyms", gymService.GetAllGyms).Methods("GET")
	router.HandleFunc(base_url+"/gym", gymService.CreateGym).Methods("POST")
	router.HandleFunc(base_url+"/gym/{gym_id}", gymService.GetGym).Methods("GET")

	router.HandleFunc(base_url+"/outcome", outcomeService.CreateOutcome).Methods("POST")
	router.HandleFunc(base_url+"/outcome/{outcome_id}", outcomeService.GetOutcome).Methods("GET")
	router.HandleFunc(base_url+"/outcome/bout/{bout_id}", outcomeService.GetOutcomeByBout).Methods("GET")
	router.HandleFunc(base_url+"/outcome/bout/{bout_id}", outcomeService.CreateOutcomeByBout).Methods("POST")

	router.HandleFunc(base_url+"/styles", styleService.GetAllStyles).Methods("GET")
	router.HandleFunc(base_url+"/style", styleService.CreateStyle).Methods("POST")
	router.HandleFunc(base_url+"/style/athlete/{athlete_id}", styleService.RegisterAthleteToStyle).Methods("POST")
	router.HandleFunc(base_url+"/styles/athlete/{athlete_id}", styleService.RegisterMultipleStylesToAthlete).Methods("POST")
	router.HandleFunc(base_url+"/styles/common/{athlete_id}/{challenger_id}", styleService.GetCommonStyles).Methods("GET")

	router.HandleFunc(base_url+"/score/{athlete_id}", athleteScoreService.GetAthleteScore).Methods("GET")
	router.HandleFunc(base_url+"/score/{athlete_id}/all", athleteScoreService.GetAllAthleteScoresByAthleteId).Methods("GET")
	router.HandleFunc(base_url+"/score/{athlete_id}/style/{style_id}", athleteScoreService.GetAthleteScoreByStyle).Methods("GET")

	router.HandleFunc(base_url+"/feed/{athlete_id}", feedService.GetFeedByAthleteId).Methods("GET")

	return router
}
