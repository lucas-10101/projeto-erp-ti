package dao

import (
	"erp/organization-api/data/database"
	"erp/organization-api/data/entities"
)

type UserCompanyDAO struct {
	Connection database.Connection
}

func (dao *UserCompanyDAO) Create(uc *entities.UserCompany) error {
	_, err := dao.Connection.Exec(
		"INSERT INTO user_companies (user_id, company_id) VALUES ($1, $2)",
		uc.UserId, uc.CompanyId,
	)
	return err
}

func (dao *UserCompanyDAO) Delete(userId, companyId int64) error {
	_, err := dao.Connection.Exec(
		"DELETE FROM user_companies WHERE user_id = $1 AND company_id = $2",
		userId, companyId,
	)
	return err
}
