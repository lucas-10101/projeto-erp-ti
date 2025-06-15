package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type AccessGroupRoleDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *AccessGroupRoleDAO) Create(agr *entities.AccessGroupRole) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"INSERT INTO access_group_roles (access_group_id, role_id) VALUES (?1, ?2)",
		agr.AccessGroupId, agr.RoleId,
	)
	return err
}

func (dao *AccessGroupRoleDAO) Read(accessGroupId int64, roleId string) (*entities.AccessGroupRole, error) {
	agr := &entities.AccessGroupRole{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		"SELECT access_group_id, role_id FROM access_group_roles WHERE access_group_id = ?1 AND role_id = ?2",
		accessGroupId, roleId,
	).Scan(&agr.AccessGroupId, &agr.RoleId)
	return agr, err
}

func (dao *AccessGroupRoleDAO) Delete(accessGroupId int64, roleId string) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM access_group_roles WHERE access_group_id = ?1 AND role_id = ?2",
		accessGroupId, roleId,
	)
	return err
}
