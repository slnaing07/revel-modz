package rbac

type RoleInterface interface {
}

type Role struct {
	RoleId   int64 // primary key
	RoleName string
}

type RoleHierarchyInterface interface {
}

type RoleHierarchy struct {
	RoleId int64

	ParentRole int64
	RoleName   string
}

type AdminRoleInterface interface {
}

type AdminRole struct {
	AdminRoleId   int64
	AdminRoleName string
}

type AdminRoleHierarchyInterface interface {
}

type AdminRoleHierarchy struct {
	AdminRoleId int64

	AdminParentRole int64
	AdminRoleName   string
}
