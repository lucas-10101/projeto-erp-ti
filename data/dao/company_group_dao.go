package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type CompanyGroupDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *CompanyGroupDAO) Create(cg *entities.CompanyGroup) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`INSERT INTO company_groups 
		(id, name, activate, country_id, country_subdivision_id) 
		VALUES ($1, $2, $3, $4, $5)`,
		cg.Id, cg.Name, cg.Activate, cg.CountryId, cg.CountrySubdivisionId,
	)
	return err
}

func (dao *CompanyGroupDAO) Read(id int64) (*entities.CompanyGroup, error) {
	cg := &entities.CompanyGroup{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		`SELECT id, name, activate, country_id, country_subdivision_id 
		FROM company_groups WHERE id = $1`,
		id,
	).Scan(
		&cg.Id, &cg.Name, &cg.Activate, &cg.CountryId, &cg.CountrySubdivisionId,
	)
	return cg, err
}

func (dao *CompanyGroupDAO) Update(cg *entities.CompanyGroup) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`UPDATE company_groups SET 
		name = $1, activate = $2, country_id = $3, country_subdivision_id = $4 
		WHERE id = $5`,
		cg.Name, cg.Activate, cg.CountryId, cg.CountrySubdivisionId, cg.Id,
	)
	return err
}

func (dao *CompanyGroupDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		"DELETE FROM company_groups WHERE id = $1",
		id,
	)
	return err
}
