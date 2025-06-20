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

func TestCreateCompanyGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
	}

	id, err := dao.Create(cg)
	if err != nil {
		t.Errorf("Error at creation: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id not valid:")
	}
}

func TestCreateCompanyGroupWithNonNullValues(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	someId := int64(1234)
	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Id:                   nil,
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            &someId,
		CountrySubdivisionId: &someId,
	}

	id, err := dao.Create(cg)
	if err != nil {
		t.Errorf("Error at creation: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id not valid:")
	}
}

func TestReadCompanyGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Id:                   nil,
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
	}

	id, err := dao.Create(cg)
	if err != nil {
		t.Errorf("Error at creation: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id not valid:")
	}

	cg.Id = &id

	var returned *entities.CompanyGroup
	returned, err = dao.Read(id)

	if err != nil || returned == nil {
		t.Errorf("Error at reading: %v", err)
	}

	var returnedJsonBytes, expectedJsonBytes []byte
	returnedJsonBytes, err = json.Marshal(returned)
	expectedJsonBytes, err = json.Marshal(cg)

	if string(returnedJsonBytes) != string(expectedJsonBytes) {
		t.Errorf("Returned JSON does not match expected JSON:\nExpected: %s\nReturned: %s", string(returnedJsonBytes), string(expectedJsonBytes))
	}
}

func TestUpdateCompanyGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Id:                   nil,
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
	}

	id, err := dao.Create(cg)
	if err != nil {
		t.Errorf("Error at creation: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id not valid:")
	}

	cg.Id = &id

	cg.Name = "Updated Group Name"

	if err = dao.Update(cg); err != nil {
		t.Errorf("Error at update: %v", err)
	}

	var returned *entities.CompanyGroup
	returned, err = dao.Read(id)

	if err != nil || returned == nil {
		t.Errorf("Error at reading: %v", err)
	}

	var returnedJsonBytes, expectedJsonBytes []byte
	returnedJsonBytes, err = json.Marshal(returned)
	expectedJsonBytes, err = json.Marshal(cg)

	if string(returnedJsonBytes) != string(expectedJsonBytes) {
		t.Errorf("Returned JSON does not match expected JSON:\nExpected: %s\nReturned: %s", string(returnedJsonBytes), string(expectedJsonBytes))
	}
}

func TestDeleteCompanyGroup(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Id:                   nil,
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
	}

	id, err := dao.Create(cg)
	if err != nil {
		t.Errorf("Error at creation: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id not valid:")
	}

	if err = dao.Delete(id); err != nil {
		t.Errorf("Error at deletion: %v", err)
	}

	_, err = dao.Read(id)

	if err != sql.ErrNoRows {
		t.Errorf("Expected no rows error after deletion, got: %v", err)
	}
}
