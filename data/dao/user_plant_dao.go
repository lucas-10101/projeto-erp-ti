package dao

import (
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type UserPlantDAO struct {
	Connection database.Connection
}

func (dao *UserPlantDAO) Create(up *entities.UserPlant) error {
	_, err := dao.Connection.Exec(
		"INSERT INTO user_plants (user_id, plant_id) VALUES ($1, $2)",
		up.UserId, up.PlantId,
	)
	return err
}

func (dao *UserPlantDAO) Delete(userId, plantId int64) error {
	_, err := dao.Connection.Exec(
		"DELETE FROM user_plants WHERE user_id = $1 AND plant_id = $2",
		userId, plantId,
	)
	return err
}
