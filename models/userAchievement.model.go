package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"          json:"_id"`
	UserID       string             `bson:"user_id,omitempty"      json:"user_id"`
	Achievements []UserAchievement  `bson:"achievements,omitempty" json:"achievements"`
	CreatedAt    primitive.DateTime `bson:"createdAt,omitempty"    json:"createdAt"`
	UpdatedAt    primitive.DateTime `bson:"updatedAt,omitempty"    json:"updatedAt"`
}

type UserAchievement struct {
	AchievementID primitive.ObjectID `bson:"achievement_id,omitempty" json:"achievement_id"`
	UserValue     float64            `bson:"user_value"               json:"user_value"`
	Reached       bool               `bson:"reached"                  json:"reached"`
	CreatedAt     primitive.DateTime `bson:"createdAt,omitempty"      json:"createdAt"`
	UpdatedAt     primitive.DateTime `bson:"updatedAt,omitempty"      json:"updatedAt"`
}
