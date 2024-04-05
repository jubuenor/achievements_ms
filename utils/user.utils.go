package utils

type NewUser struct {
	UserID string `json:"user_id"`
}

type UserAchievementValueUpdate struct {
	UserID        string  `json:"user_id"`
	AchievementID string  `json:"achievement_id"`
	Value         float64 `json:"value"`
}
