package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	models "github.com/hizbi-github/gost/new-project-core/models"
)

func Get(ctx context.Context, database *mongo.Database, someId string) error {
	_, err := database.Collection("some_collection").Find(ctx, bson.D{{Key: "some_id", Value: someId}})
	if err != nil {
		return err
	}

	return err
}

func Exists(ctx context.Context, database *mongo.Database, someId string) bool {
	err := database.Collection("some_collection").FindOne(ctx, bson.D{{Key: "some_id", Value: someId}}).Err()
	if err != nil && err == mongo.ErrNoDocuments {
		return false
	} else {
		return true
	}
}

func Save(ctx context.Context, database *mongo.Database, someMongoDocument *models.SomeMongoDocument) error {
	_, err := database.Collection("some_collection").InsertOne(ctx, someMongoDocument)
	return err
}
