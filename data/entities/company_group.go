package entities

type CompanyGroup struct {
	Id                   int64
	Name                 string
	Activate             bool
	CountryId            int64
	CountrySubdivisionId int64
}
