package container

import (
	"errors"
)

type Container interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	Clear() error
	Keys() ([]string, error)
	Values() ([]interface{}, error)
	Len() (int, error)
	Has(key string) (bool, error)
	IsEmpty() (bool, error)
}

var (
	errNoElements      = errors.New("no elements")
	errIndexOutOfRange = errors.New("index out of range")
	errNotInitialized  = errors.New("container is initialized")
)
