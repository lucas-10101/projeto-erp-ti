package main

import (
	"erp/organization-api/data/database"
	"erp/organization-api/utils"
	"fmt"
)

func Init() {
	database.CreateConnection() // Initialize the database connection
}

func main() {

	if err := utils.LoadApplicationPropertiesFromFile("application.properties"); err != nil {
		fmt.Println("Error loading properties:", err)
	}

	Init()
}
