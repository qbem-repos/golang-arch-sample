package repository

import "github.com/qbem-repository/golang-arch-sample/internal/foo/entity"

type FooRepository interface {
	Add(etty *entity.Foo) error
}
