package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
	"erp/organization-api/utils"
	"testing"
)

func TestCreateAccessGroupRole(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	accessGroupDAO := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	idAccessGroupDAO, err := accessGroupDAO.Create(accessGroup)

	if err != nil || idAccessGroupDAO <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}

	roleDAO := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	roleId, err := roleDAO.Create(role)

	if err != nil || roleId <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}

	dao := &AccessGroupRoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	agr := &entities.AccessGroupRole{
		AccessGroupId: idAccessGroupDAO,
		RoleId:        roleId,
	}

	err = dao.Create(agr)
	if err != nil {
		t.Errorf("Failed to create AccessGroupRole: %v", err)
	}
}

func TestCreateAccessGroupRole2(t *testing.T) {

	t.Setenv("APPLICATION_PROPERTIES_FILE", "../../application.properties")

	utils.LoadApplicationPropertiesFromFile()
	database.CreateConnection()

	accessGroupDAO := &AccessGroupDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	accessGroup := &entities.AccessGroup{
		Name: "Test AccessGroup",
	}

	idAccessGroupDAO, err := accessGroupDAO.Create(accessGroup)

	if err != nil || idAccessGroupDAO <= 0 {
		t.Fatalf("Failed to create accessGroup: %v", err)
	}

	roleDAO := &RoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	role := &entities.Role{
		Name: "Test Role",
	}

	roleId, err := roleDAO.Create(role)

	if err != nil || roleId <= 0 {
		t.Fatalf("Failed to create role: %v", err)
	}

	dao := &AccessGroupRoleDAO{
		Connection: database.GetConnection(),
		Ctx:        context.Background(),
	}

	agr := &entities.AccessGroupRole{
		AccessGroupId: idAccessGroupDAO,
		RoleId:        roleId,
	}

	err = dao.Create(agr)
	if err != nil {
		t.Errorf("Failed to create AccessGroupRole: %v", err)
	}

	// Test Exists
	exists := dao.Exists(idAccessGroupDAO, roleId)
	if !exists {
		t.Errorf("AccessGroupRole should exist after creation")
	}
	// Test Delete
	err = dao.Delete(idAccessGroupDAO, roleId)
	if err != nil {
		t.Errorf("Failed to delete AccessGroupRole: %v", err)
	}
}
