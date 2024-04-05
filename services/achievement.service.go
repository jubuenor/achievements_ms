package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FinanceUN/Achievements/db"
	"github.com/FinanceUN/Achievements/models"
)

func CreateAchievement(achievement models.Achievement) (*mongo.InsertOneResult, error) {
	newAchievement := models.Achievement{
		Title:            achievement.Title,
		Description:      achievement.Description,
		RequirementValue: achievement.RequirementValue,
		AchievementTier:  achievement.AchievementTier,
		CreatedAt:        primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:        primitive.NewDateTimeFromTime(time.Now()),
	}

	if achievement.NextAchievement != primitive.NilObjectID {
		newAchievement.NextAchievement = achievement.NextAchievement
	} else {
		newAchievement.NextAchievement = primitive.NilObjectID
	}

	insertResult, err := db.AchievementsCollection.InsertOne(context.Background(), newAchievement)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func GetAchievements() ([]models.Achievement, error) {
	var achievements []models.Achievement

	cursor, err := db.AchievementsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var achievement models.Achievement
		cursor.Decode(&achievement)
		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

func GetAchievementsByTier(tier int) ([]models.Achievement, error) {
	var achievements []models.Achievement

	cursor, err := db.AchievementsCollection.Find(
		context.Background(),
		bson.D{{"achievementTier", tier}},
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var achievement models.Achievement
		cursor.Decode(&achievement)
		achievements = append(achievements, achievement)
	}

	return achievements, nil
}

func GetAchievement(id string) (*models.Achievement, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var achievement models.Achievement
	err = db.AchievementsCollection.FindOne(context.Background(), bson.D{{"_id", objectID}}).
		Decode(&achievement)
	if err != nil {
		return nil, err
	}

	return &achievement, nil
}

func UpdateAchievement(achievement models.Achievement) (*mongo.UpdateResult, error) {
	achievement.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	updateResult, err := db.AchievementsCollection.UpdateOne(context.Background(),
		bson.D{{"_id", achievement.ID}},
		bson.D{{"$set", achievement}})
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func DeleteAchievement(id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	deleteResult, err := db.AchievementsCollection.DeleteOne(
		context.Background(),
		bson.D{{"_id", objectID}},
	)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
