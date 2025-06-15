package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type UserAccessGroupDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *UserAccessGroupDAO) Create(uag *entities.UserAccessGroup) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"INSERT INTO user_access_groups (user_id, access_group_id) VALUES ($1, $2)",
		uag.UserId, uag.AccessGroupId,
	)
	return err
}

func (dao *UserAccessGroupDAO) Delete(userId, accessGroupId int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM user_access_groups WHERE user_id = $1 AND access_group_id = $2",
		userId, accessGroupId,
	)
	return err
}
