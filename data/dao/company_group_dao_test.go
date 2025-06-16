package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
	"erp/organization-api/utils"
	"testing"
)

func TestCreate(t *testing.T) {

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

func TestCreateWithNonNullValues(t *testing.T) {

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
