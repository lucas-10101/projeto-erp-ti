package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type AccessGroupDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *AccessGroupDAO) Create(ag *entities.AccessGroup) (id int64, err error) {
	row := dao.Connection.QueryRowContext(
		dao.Ctx,
		"INSERT INTO access_groups (name) VALUES ($1) returning id",
		ag.Name,
	)

	if err = row.Err(); err == nil {
		err = row.Scan(&id)
	}

	return id, err
}

func (dao *AccessGroupDAO) Read(id int64) (*entities.AccessGroup, error) {
	ag := &entities.AccessGroup{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		"SELECT id, name FROM access_groups WHERE id = $1",
		id,
	).Scan(&ag.Id, &ag.Name)
	return ag, err
}

func (dao *AccessGroupDAO) Update(ag *entities.AccessGroup) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"UPDATE access_groups SET name = $1 WHERE id = $2",
		ag.Name, ag.Id,
	)
	return err
}

// TODO: Remove roles before deleting
func (dao *AccessGroupDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM access_groups WHERE id = $1",
		id,
	)
	return err
}
