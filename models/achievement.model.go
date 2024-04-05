package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Achievement struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"              json:"_id"`
	Title            string             `bson:"title,omitempty"            json:"title"`
	Description      string             `bson:"description,omitempty"      json:"description"`
	RequirementValue float64            `bson:"requirementValue,omitempty" json:"requirementValue"`
	AchievementTier  int                `bson:"achievementTier,omitempty"  json:"achievementTier"`
	NextAchievement  primitive.ObjectID `bson:"nextAchievement"            json:"nextAchievement"`
	CreatedAt        primitive.DateTime `bson:"createdAt,omitempty"        json:"createdAt"`
	UpdatedAt        primitive.DateTime `bson:"updatedAt,omitempty"        json:"updatedAt"`
}
