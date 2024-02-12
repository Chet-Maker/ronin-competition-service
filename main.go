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
	// services.SetDB(dbconn)

	athleteRepo := repositories.NewAthleteRepository(dbconn)
	feedRepo := repositories.NewFeedRepository(dbconn)
	styleRepo := repositories.NewStyleRepository(dbconn)
	boutRepo := repositories.NewBoutRepository(dbconn)
	outcomeRepo := repositories.NewOutcomeRepository(dbconn)
	if outcomeRepo == nil || outcomeRepo.DB == nil {
		log.Fatal("Failed to initialize outcome repository")
	}
	athleteScoreRepo := repositories.NewAthleteScoreRepository(dbconn)

	// services.SetAthleteRepo(athleteRepo)
	// services.SetFeedRepo(feedRepo)
	// services.SetStyleRepo(styleRepo)
	// services.SetBoutRepo(boutRepo)
	// services.SetOutcomeRepo(outcomeRepo)
	// services.SetAthleteScoreRepo(athleteScoreRepo)

	//Using Constructor Injection

	athleteScoreService := services.NewAthleteScoreService(athleteScoreRepo)
	athleteService := services.NewAthleteService(athleteRepo)
	feedService := services.NewFeedService(feedRepo)
	styleService := services.NewStyleService(athleteScoreService, styleRepo)
	boutService := services.NewBoutService(boutRepo)
	outcomeService := services.NewOutcomeService(athleteScoreService, outcomeRepo, boutRepo)
	if outcomeService == nil {
		log.Fatal("Failed to initialize outcome service")
	}

	var appRouter = router.CreateRouter(dbconn, athleteService, feedService, styleService, boutService, outcomeService, athleteScoreService)

	log.Println("listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
