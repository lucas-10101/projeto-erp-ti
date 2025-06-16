package main

import (
	"erp/organization-api/data/database"
	"erp/organization-api/service"
	"erp/organization-api/utils"
	"fmt"
)

func Init() {
	database.CreateConnection()        // Initialize the database connection
	database.CreateMongoDBConnection() // Initialize the MongoDB connection
	service.InitLogger()               // Initialize the logger
}

func main() {
	if err := utils.LoadApplicationPropertiesFromFile(); err != nil {
		panic(fmt.Sprintf("Error loading properties:", err))
	}
	Init()
}
