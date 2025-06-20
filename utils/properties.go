package utils

import (
	"bufio"
	"os"
	"reflect"
	"strings"
)

var (
	ApplicationProperties = Properties{}
)

type Properties struct {
	DatabaseDriver           string
	DatabaseConnectionString string
	ApplicationStage         string
	ApplicationName          string
	LogFile                  string
	LogLevel                 int

	MongoDBConnectionString string
	MongoBDDatabaseName     string
	MongoDBCollectionName   string
}

// Load FromFile reads properties from a file named "application.properties" by reflection
func LoadApplicationPropertiesFromFile() (err error) {

	fileName := os.Getenv("APPLICATION_PROPERTIES_FILE")

	var file *os.File
	file, err = os.Open(fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	fields := reflect.VisibleFields(reflect.TypeOf(&ApplicationProperties).Elem())
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if scanner.Err() != nil {
			return scanner.Err()
		}

		line := strings.TrimSpace(scanner.Text())

		// Ignore comments
		if commentStart := strings.Index(line, "#"); commentStart != -1 {
			line = strings.TrimSpace(line[:commentStart])
		}
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if parts[0] == "" || len(parts) < 2 {
			continue
		}

		propertyName := strings.TrimSpace(parts[0])
		propertyValue := strings.TrimSpace(parts[1])

		for _, field := range fields {
			if strings.EqualFold(field.Name, propertyName) {
				fieldValue := reflect.ValueOf(&ApplicationProperties).Elem().FieldByName(field.Name)
				if fieldValue.IsValid() && fieldValue.CanSet() && fieldValue.Kind() == reflect.String {
					fieldValue.SetString(propertyValue)
				}
				break
			}
		}
	}

	return nil
}
