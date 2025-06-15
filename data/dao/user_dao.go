package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type UserDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *UserDAO) Create(user *entities.User) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`INSERT INTO users (
			id, 
			username, 
			password, 
			active
		) VALUES ($1, $2, $3, $4)`,
		user.Id,
		user.Username,
		user.Password,
		user.Active,
	)
	return err
}

func (dao *UserDAO) Read(id int64) (*entities.User, error) {
	user := &entities.User{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		`SELECT 
			id, 
			username, 
			password, 
			active 
		FROM users 
		WHERE id = $1`,
		id,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Active,
	)
	return user, err
}

func (dao *UserDAO) Update(user *entities.User) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`UPDATE users SET 
			username = $1, 
			password = $2, 
			active = $3 
		WHERE id = $4`,
		user.Username,
		user.Password,
		user.Active,
		user.Id,
	)
	return err
}

func (dao *UserDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`DELETE FROM users 
		WHERE id = $1`,
		id,
	)
	return err
}
