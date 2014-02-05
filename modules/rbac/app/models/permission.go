package rbac

type UserPermissionInterface interface {
}

type UserPermission struct {
	PermissionId int64

	OperationId    int64
	ObjectId       int64
	PermissionName string
}

type AdminPermissionInterface interface {
}

type AdminPermission struct {
	AdminPermissionId int64

	AdminOperationId    int64
	AdminObjectId       int64
	AdminPermissionName string
}
