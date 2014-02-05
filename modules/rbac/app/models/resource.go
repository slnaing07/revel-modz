package rbac

type ResourceInterface interface {
}

type Resource struct {
	ResourceId int64

	ResourceName string

	Locked bool
}
