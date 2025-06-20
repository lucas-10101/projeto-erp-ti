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

func TestCreateCompany(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	daoCompanyGroup := &CompanyGroupDAO{
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

	id, err := daoCompanyGroup.Create(cg)
	if err != nil {
		t.Errorf("Error at creation of company group: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id of company group is not valid:")
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
		CompanyGroupId:       &id,
	}

	id, err = dao.Create(company)
	if err != nil {
		t.Errorf("Error at creation of company: %v", err)
	}

	if id <= 0 {
		t.Errorf("Company id not valid:")
	}
}

func TestCreateCompanyWithoutCompany(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
	}

	_, err := dao.Create(company)
	if err == nil {
		t.Errorf("Expected error at creation, got nil")
	}
}

func TestCreateCompanyWithNonNullValues(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	daoCompanyGroup := &CompanyGroupDAO{
		Connection: database.GetConnection(), // Replace with actual mock connection
		Ctx:        context.Background(),
	}

	someId := int64(1234)
	// Create a new company group entity
	cg := &entities.CompanyGroup{
		Name:                 "Test Group",
		Activate:             true,
		CountryId:            &someId,
		CountrySubdivisionId: &someId,
	}

	id, err := daoCompanyGroup.Create(cg)
	if err != nil {
		t.Errorf("Error at creation of company group: %v", err)
	}

	if id <= 0 {
		t.Errorf("Id of company group is not valid:")
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
		CompanyGroupId:       &id,
	}

	id, err = dao.Create(company)
	if err != nil {
		t.Errorf("Error at creation of company: %v", err)
	}

	if id <= 0 {
		t.Errorf("Company id not valid:")
	}
}

func TestReadCompany(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	daoCompanyGroup := &CompanyGroupDAO{
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

	companyGroupId, err := daoCompanyGroup.Create(cg)
	if err != nil {
		t.Errorf("Error at creation of company group: %v", err)
	}

	if companyGroupId <= 0 {
		t.Errorf("Id of company group is not valid:")
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
		CompanyGroupId:       &companyGroupId,
	}

	id, err := dao.Create(company)
	if err != nil {
		t.Errorf("Error at creation of company: %v", err)
	}

	if id <= 0 {
		t.Errorf("Company id not valid:")
	}

	company.Id = &id

	var returned *entities.Company
	returned, err = dao.Read(id)

	if err != nil || returned == nil {
		t.Errorf("Error at reading: %v", err)
	}

	var returnedJsonBytes, expectedJsonBytes []byte
	returnedJsonBytes, err = json.Marshal(returned)
	expectedJsonBytes, err = json.Marshal(company)

	if string(returnedJsonBytes) != string(expectedJsonBytes) {
		t.Errorf("Returned JSON does not match expected JSON:\nExpected: %s\nReturned: %s", string(returnedJsonBytes), string(expectedJsonBytes))
	}
}

func TestUpdateCompany(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	daoCompanyGroup := &CompanyGroupDAO{
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

	companyGroupId, err := daoCompanyGroup.Create(cg)
	if err != nil {
		t.Errorf("Error at creation of company group: %v", err)
	}

	if companyGroupId <= 0 {
		t.Errorf("Id of company group is not valid:")
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
		CompanyGroupId:       &companyGroupId,
	}

	id, err := dao.Create(company)
	if err != nil {
		t.Errorf("Error at creation of company: %v", err)
	}

	if id <= 0 {
		t.Errorf("Company id not valid:")
	}

	company.Id = &id

	company.Name = "Updated Company Name"

	if err = dao.Update(company); err != nil {
		t.Errorf("Error at update: %v", err)
	}

	var returned *entities.Company
	returned, err = dao.Read(id)

	if err != nil || returned == nil {
		t.Errorf("Error at reading: %v", err)
	}

	var returnedJsonBytes, expectedJsonBytes []byte
	returnedJsonBytes, err = json.Marshal(returned)
	expectedJsonBytes, err = json.Marshal(company)

	if string(returnedJsonBytes) != string(expectedJsonBytes) {
		t.Errorf("Returned JSON does not match expected JSON:\nExpected: %s\nReturned: %s", string(returnedJsonBytes), string(expectedJsonBytes))
	}
}

func TestDeleteCompany(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	dao := &CompanyDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	daoCompanyGroup := &CompanyGroupDAO{
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

	companyGroupId, err := daoCompanyGroup.Create(cg)
	if err != nil {
		t.Errorf("Error at creation of company group: %v", err)
	}

	if companyGroupId <= 0 {
		t.Errorf("Id of company group is not valid:")
	}

	// Create a new company  entity
	company := &entities.Company{
		Name:                 "Test ",
		Activate:             true,
		CountryId:            nil,
		CountrySubdivisionId: nil,
		CompanyGroupId:       &companyGroupId,
	}

	id, err := dao.Create(company)
	if err != nil {
		t.Errorf("Error at creation of company: %v", err)
	}

	if id <= 0 {
		t.Errorf("Company id not valid:")
	}

	if err = dao.Delete(id); err != nil {
		t.Errorf("Error at deletion: %v", err)
	}

	_, err = dao.Read(id)

	if err != sql.ErrNoRows {
		t.Errorf("Expected no rows error after deletion, got: %v", err)
	}
}
