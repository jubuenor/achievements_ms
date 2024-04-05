package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/models"
	"github.com/FinanceUN/Achievements/services"
)

func CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	json.NewDecoder(r.Body).Decode(&achievement)

	result, err := services.CreateAchievement(achievement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func GetAchievements(w http.ResponseWriter, r *http.Request) {
	result, err := services.GetAchievements()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAchievementsByTier(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tier, err := strconv.Atoi(params["tier"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	result, err := services.GetAchievementsByTier(tier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAchievement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := services.GetAchievement(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	json.NewDecoder(r.Body).Decode(&achievement)

	result, err := services.UpdateAchievement(achievement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := services.DeleteAchievement(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
