package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type RoleDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *RoleDAO) Create(role *entities.Role) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`INSERT INTO roles (
			id, 
			name
		) VALUES (?1, ?2)`,
		role.Id,
		role.Name,
	)
	return err
}

func (dao *RoleDAO) Read(id int64) (*entities.Role, error) {
	role := &entities.Role{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		`SELECT 
			id, 
			name 
		FROM roles 
		WHERE id = ?1`,
		id,
	).Scan(
		&role.Id,
		&role.Name,
	)
	return role, err
}

func (dao *RoleDAO) Update(role *entities.Role) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`UPDATE roles SET 
			name = ?1 
		WHERE id = ?2`,
		role.Name,
		role.Id,
	)
	return err
}

func (dao *RoleDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`DELETE FROM roles 
		WHERE id = ?1`,
		id,
	)
	return err
}
