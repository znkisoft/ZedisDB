package container

type Container interface {
	Create(name string) error
	Drop(name string) error
}
