package service

import (
	"erp/organization-api/data/database"
	"erp/organization-api/utils"
	"log/slog"
)

var (
	Logger *slog.Logger
)

func InitLogger() {
	slog.SetDefault(slog.New(GetMongoDBLogHandler()))
}

func GetMongoDBLogHandler() *utils.MongoDBLogHandler {

	mongodbClient := database.GetMongoDBConnection()

	handler := utils.MongoDBLogHandler{
		ApplicationName: utils.ApplicationProperties.ApplicationName,
		DatabaseName:    utils.ApplicationProperties.MongoBDDatabaseName,
		CollectionName:  utils.ApplicationProperties.MongoDBCollectionName,
		Client:          mongodbClient,
		Level:           slog.Level(utils.ApplicationProperties.LogLevel),
	}

	return &handler
}
