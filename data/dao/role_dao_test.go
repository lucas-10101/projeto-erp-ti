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

func TestCreateRole(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	id, err := dao.Create(role)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}
}

func TestReadRole(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	id, err := dao.Create(role)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}

	role.Id = &id

	var found *entities.Role
	found, err = dao.Read(id)

	if err != nil {
		t.Fatalf("Failed to read role: %v", err)
	}

	expected, err := json.Marshal(role)
	if err != nil {
		t.Fatalf("Failed to marshal expected role: %v", err)
	}

	actual, err := json.Marshal(found)
	if err != nil {
		t.Fatalf("Failed to marshal found role: %v", err)
	}

	if string(expected) != string(actual) {
		t.Errorf("Expected role %s, but got %s", expected, actual)
	}

}

func TestUpdateRole(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	id, err := dao.Create(role)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}

	role.Id = &id
	role.Name = "Updated Role"

	err = dao.Update(role)
	if err != nil {
		t.Fatalf("Failed to update role: %v", err)
	}

	var found *entities.Role
	found, err = dao.Read(id)

	if err != nil {
		t.Fatalf("Failed to read role: %v", err)
	}

	expected, err := json.Marshal(role)
	if err != nil {
		t.Fatalf("Failed to marshal expected role: %v", err)
	}

	actual, err := json.Marshal(found)
	if err != nil {
		t.Fatalf("Failed to marshal found role: %v", err)
	}

	if string(expected) != string(actual) {
		t.Errorf("Expected role %s, but got %s", expected, actual)
	}

}

func TestDeketeRole(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	id, err := dao.Create(role)

	if err != nil || id <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}

	if err = dao.Delete(id); err != nil {
		t.Fatalf("Failed to delete role: %v", err)
	}

	_, err = dao.Read(id)

	if err != sql.ErrNoRows {
		t.Fatalf("Expected no rows, but got error: %v", err)
	}
}
