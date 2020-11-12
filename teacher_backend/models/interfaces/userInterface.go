package interfaces

type User interface {
	GetIdentity() string
	GetName() string
	GetType() int
}
