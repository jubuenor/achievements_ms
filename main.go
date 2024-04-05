package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/db"
	"github.com/FinanceUN/Achievements/routes"
)

func main() {
	fmt.Println("Starting Achievements server...")

	client := db.DBConnection()

	router := mux.NewRouter()

	routes.IndexRoutes(router)
	routes.AchievementsRoutes(router)
	routes.UserAchievementRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server started on port " + port + "...")
	http.ListenAndServe(":"+port, router)

	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
