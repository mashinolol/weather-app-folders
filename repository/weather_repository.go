package repository

import (
	"context"

	"weather-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeatherRepository struct {
	collection *mongo.Collection
}

func NewWeatherRepository(client *mongo.Client, dbName string) *WeatherRepository {
	return &WeatherRepository{
		collection: client.Database(dbName).Collection("weather"),
	}
}

func (r *WeatherRepository) GetWeather(ctx context.Context, city string) (*models.WeatherData, error) {
	filter := bson.M{"city": city}
	var weather models.WeatherData
	err := r.collection.FindOne(ctx, filter).Decode(&weather)
	return &weather, err
}

func (r *WeatherRepository) UpdateWeather(ctx context.Context, weather *models.WeatherData) error {
	filter := bson.M{"city": weather.City}
	update := bson.M{"$set": weather}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
