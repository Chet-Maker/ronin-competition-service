package main

import (
	"log"
	"net/http"

	"ronin/repositories"
	"ronin/router"
	"ronin/services"
	"ronin/utils"
)

func main() {
	log.Println("In Main App")

	dbconn := utils.GetConnection()

	athleteRepo := repositories.NewAthleteRepository(dbconn)
	feedRepo := repositories.NewFeedRepository(dbconn)
	styleRepo := repositories.NewStyleRepository(dbconn)
	boutRepo := repositories.NewBoutRepository(dbconn)
	outcomeRepo := repositories.NewOutcomeRepository(dbconn)
	athleteScoreRepo := repositories.NewAthleteScoreRepository(dbconn)
	gymRepo := repositories.NewGymRepository(dbconn)

	//Using Constructor Injection

	athleteScoreService := services.NewAthleteScoreService(athleteScoreRepo)
	athleteService := services.NewAthleteService(athleteRepo)
	feedService := services.NewFeedService(feedRepo)
	styleService := services.NewStyleService(athleteScoreService, styleRepo)
	boutService := services.NewBoutService(boutRepo)
	outcomeService := services.NewOutcomeService(athleteScoreService, outcomeRepo, boutRepo)
	gymService := services.NewGymService(gymRepo)

	var appRouter = router.CreateRouter(dbconn, athleteService, feedService, styleService, boutService, outcomeService, athleteScoreService, gymService)

	log.Println("listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
