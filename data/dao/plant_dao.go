package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type PlantDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *PlantDAO) Create(plant *entities.Plant) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"INSERT INTO plants (id, name, activate) VALUES ($1, $2, $3)",
		plant.Id, plant.Name, plant.Activate,
	)
	return err
}

func (dao *PlantDAO) Read(id int64) (*entities.Plant, error) {
	plant := &entities.Plant{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		"SELECT id, name, activate FROM plants WHERE id = $1",
		id,
	).Scan(&plant.Id, &plant.Name, &plant.Activate)
	return plant, err
}

func (dao *PlantDAO) Update(plant *entities.Plant) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"UPDATE plants SET name = $1, activate = $2 WHERE id = $3",
		plant.Name, plant.Activate, plant.Id,
	)
	return err
}

func (dao *PlantDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM plants WHERE id = $1",
		id,
	)
	return err
}
