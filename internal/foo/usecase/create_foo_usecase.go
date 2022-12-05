package usecase

import (
	"github.com/qbem-repository/golang-arch-sample/internal/foo/repository"
)

func CreateFooUsecase(input CreateFooInbound, repo repository.FooRepository) (CreateFooOutbound, error) {
	repo.Add(input.Entity())
}
