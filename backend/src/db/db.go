package db

import (
	"log"

	"github.com/chrispeterjeyaraj/track-my-wealth/server/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func Init() *gorm.DB {
	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection = client.Database("synonyms").Collection("tweets")
	return collection.Find(ctx, bson.D{})
}