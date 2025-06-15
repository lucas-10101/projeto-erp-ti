package entities

type User struct {
	Id       int64
	Username string
	Password string
	Active   bool
}

type UserCompany struct {
	UserId    int64
	CompanyId int64
}

type UserPlant struct {
	UserId  int64
	PlantId int64
}

type UserAccessGroup struct {
	UserId        int64
	AccessGroupId int64
}
