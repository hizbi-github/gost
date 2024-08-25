package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	moduleHandler "github.com/hizbi-github/gost/new-project-core/module/handler"
	moduleUsecase "github.com/hizbi-github/gost/new-project-core/module/usecase"
	mongoConnector "github.com/hizbi-github/gost/new-project-core/service/db"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	hostAndPort := "0.0.0.0:8000"
	if os.Getenv("CORE_PORT") != "" {
		hostAndPort = fmt.Sprintf("0.0.0.0:%s", os.Getenv("CORE_PORT"))
	}

	dbClient := mongoConnector.ConnectToDatabase()
	defer mongoConnector.CloseDatabaseClient(dbClient)

	database := dbClient.Database(os.Getenv("DB_DATABASE"))

	//_, err := database.Collection("original_articles").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
	//	Keys:    bson.D{{Key: "some_key", Value: 1}},
	//	Options: options.Index().SetUnique(true),
	//})

	//if err != nil {
	//	logrus.Fatalln(err)
	//}

	e := echo.New()
	e.GET("/v1/api/someEndpoint", moduleHandler.SomeHandler)

	go moduleUsecase.StartCronJob(database)

	logrus.Infoln("Starting core service...")
	logrus.Fatalln(e.Start(hostAndPort))
}
