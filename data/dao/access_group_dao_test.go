package dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
	"erp/organization-api/utils"
	"testing"
)

func TestCreateAccessGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	id, err := dao.Create(accessGroup)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}
}

func TestReadAccessGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	id, err := dao.Create(accessGroup)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}

	accessGroup.Id = &id

	var found *entities.AccessGroup
	found, err = dao.Read(id)

	if err != nil {
		t.Fatalf("Failed to read accessGroup: %v", err)
	}

	expected, err := json.Marshal(accessGroup)
	if err != nil {
		t.Fatalf("Failed to marshal expected accessGroup: %v", err)
	}

	actual, err := json.Marshal(found)
	if err != nil {
		t.Fatalf("Failed to marshal found accessGroup: %v", err)
	}

	if string(expected) != string(actual) {
		t.Errorf("Expected accessGroup %s, but got %s", expected, actual)
	}

}

func TestUpdateAccessGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	id, err := dao.Create(accessGroup)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}

	accessGroup.Id = &id
	accessGroup.Name = "Updated AccessGroup"

	err = dao.Update(accessGroup)
	if err != nil {
		t.Fatalf("Failed to update accessGroup: %v", err)
	}

	var found *entities.AccessGroup
	found, err = dao.Read(id)

	if err != nil {
		t.Fatalf("Failed to read accessGroup: %v", err)
	}

	expected, err := json.Marshal(accessGroup)
	if err != nil {
		t.Fatalf("Failed to marshal expected accessGroup: %v", err)
	}

	actual, err := json.Marshal(found)
	if err != nil {
		t.Fatalf("Failed to marshal found accessGroup: %v", err)
	}

	if string(expected) != string(actual) {
		t.Errorf("Expected accessGroup %s, but got %s", expected, actual)
	}

}

func TestDeleteAccessGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	id, err := dao.Create(accessGroup)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}

	if err = dao.Delete(id); err != nil {
		t.Fatalf("Failed to delete accessGroup: %v", err)
	}

	_, err = dao.Read(id)

	if err != sql.ErrNoRows {
		t.Fatalf("Expected no rows, but got error: %v", err)
	}
}
