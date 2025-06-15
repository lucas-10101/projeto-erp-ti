package entities

type Company struct {
	Id                   int64
	Name                 string
	Activate             bool
	CountryId            int64
	CountrySubdivisionId int64

	CompanyGroupId int64
}
