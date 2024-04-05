package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FinanceUN/Achievements/services"
	"github.com/FinanceUN/Achievements/utils"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user utils.NewUser
	json.NewDecoder(r.Body).Decode(&user)

	result, err := services.RegisterNewUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func UpdateAchievementUserValue(w http.ResponseWriter, r *http.Request) {
	var userUpdate utils.UserAchievementValueUpdate
	json.NewDecoder(r.Body).Decode(&userUpdate)

	result, err := services.UpdateAchievementUserValue(userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetUserAchievements(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	userID := v.Get("user_id")
	options := v.Get("options")

	result, err := services.GetUserAchievements(userID, options)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := services.DeleteUser(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
