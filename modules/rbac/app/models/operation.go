package rbac

type OperationInterface interface {
}

const (
	CRUD_CREATE = 1
	CRUD_READ   = 2
	CRUD_UPDATE = 4
	CRUD_DELETE = 8
)

type Operation struct {
	OperationId int64

	// CRUD bit mask
	// 1  create
	// 2  read
	// 4  update
	// 8  delete
	OperationCRUD byte
	OperationName string

	Locked bool
}
