package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() *mongo.Client {

	username := url.QueryEscape(os.Getenv("DB_USERNAME"))
	password := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	hostname := os.Getenv("DB_HOSTNAME")

	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s", username, password, hostname)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func CloseDatabaseClient(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
}
