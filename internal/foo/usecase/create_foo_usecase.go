package usecase

import (
	"github.com/qbem-repository/golang-arch-sample/internal/foo/entity"
	"github.com/qbem-repository/golang-arch-sample/internal/foo/repository"
)

type CreateFooInbound struct {
}

func (i *CreateFooInbound) Entity() *entity.Foo {
	return &entity.Foo{}
}

type CreateFooOutbound struct {
}

func CreateFooUsecase(input CreateFooInbound, repo repository.FooRepository) (CreateFooOutbound, error) {
	repo.Add(input.Entity())
}
