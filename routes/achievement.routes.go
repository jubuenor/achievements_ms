package routes

import (
	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/controllers"
)

func AchievementsRoutes(router *mux.Router) {
	router.HandleFunc("/achievements", controllers.CreateAchievement).Methods("POST")
	router.HandleFunc("/achievements", controllers.GetAchievements).Methods("GET")
	router.HandleFunc("/achievements/{id}", controllers.GetAchievement).Methods("GET")
	router.HandleFunc("/achievements/tier/{tier}", controllers.GetAchievementsByTier).Methods("GET")
	router.HandleFunc("/achievements", controllers.UpdateAchievement).Methods("PUT")
	router.HandleFunc("/achievements/{id}", controllers.DeleteAchievement).Methods("DELETE")
}
