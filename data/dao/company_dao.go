package dao

import (
	"context"
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type CompanyDAO struct {
	Connection database.Connection
	Ctx        context.Context
}

func (dao *CompanyDAO) Create(company *entities.Company) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`INSERT INTO companies (
			id, 
			name, 
			activate, 
			country_id, 
			country_subdivision_id, 
			company_group_id
		) VALUES ($1, $2, $3, $4, $5, $6)`,
		company.Id,
		company.Name,
		company.Activate,
		company.CountryId,
		company.CountrySubdivisionId,
		company.CompanyGroupId,
	)
	return err
}

func (dao *CompanyDAO) Read(id int64) (*entities.Company, error) {
	company := &entities.Company{}
	err := dao.Connection.QueryRowContext(
		dao.Ctx,
		`SELECT 
			id, 
			name, 
			activate, 
			country_id, 
			country_subdivision_id, 
			company_group_id 
		FROM companies 
		WHERE id = $1`,
		id,
	).Scan(
		&company.Id,
		&company.Name,
		&company.Activate,
		&company.CountryId,
		&company.CountrySubdivisionId,
		&company.CompanyGroupId,
	)
	return company, err
}

func (dao *CompanyDAO) Update(company *entities.Company) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`UPDATE companies SET 
			name = $1, 
			activate = $2, 
			country_id = $3, 
			country_subdivision_id = $4, 
			company_group_id = $5 
		WHERE id = $6`,
		company.Name,
		company.Activate,
		company.CountryId,
		company.CountrySubdivisionId,
		company.CompanyGroupId,
		company.Id,
	)
	return err
}

func (dao *CompanyDAO) Delete(id int64) error {
	_, err := dao.Connection.ExecContext(
		dao.Ctx,
		`DELETE FROM companies WHERE id = $1`,
		id,
	)
	return err
}
