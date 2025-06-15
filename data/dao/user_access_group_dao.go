package dao

import (
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type UserAccessGroupDAO struct {
	Connection database.Connection
}

func (dao *UserAccessGroupDAO) Create(uag *entities.UserAccessGroup) error {
	_, err := dao.Connection.Exec(
		"INSERT INTO user_access_groups (user_id, access_group_id) VALUES ($1, $2)",
		uag.UserId, uag.AccessGroupId,
	)
	return err
}

func (dao *UserAccessGroupDAO) Delete(userId, accessGroupId int64) error {
	_, err := dao.Connection.Exec(
		"DELETE FROM user_access_groups WHERE user_id = $1 AND access_group_id = $2",
		userId, accessGroupId,
	)
	return err
}
