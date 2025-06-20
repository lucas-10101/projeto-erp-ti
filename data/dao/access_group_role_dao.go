package dao

import (
	"context"
	"database/sql"
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
		"INSERT INTO access_group_roles (access_group_id, role_id) VALUES ($1, $2)",
		agr.AccessGroupId, agr.RoleId,
	)
	return err
}

func (dao *AccessGroupRoleDAO) Delete(accessGroupId, roleId int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM access_group_roles WHERE access_group_id = $1 and role_id = $2",
		accessGroupId, roleId,
	)
	return err
}

func (dao *AccessGroupRoleDAO) Exists(accessGroupId, roleId int64) bool {
	row := dao.Connection.QueryRowContext(
		dao.Ctx,
		"SELECT null FROM access_group_roles WHERE access_group_id = $1 AND role_id = $2",
		accessGroupId, roleId,
	)

	return row.Err() != sql.ErrNoRows
}
