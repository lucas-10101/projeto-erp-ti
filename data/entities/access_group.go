package entities

type AccessGroup struct {
	Id   *int64
	Name string
}

type AccessGroupRole struct {
	AccessGroupId int64
	RoleId        int64
}
