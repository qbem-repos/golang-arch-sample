package usecase

import (
	"github.com/qbem-repository/golang-arch-sample/internal/foo/entity"
	"github.com/qbem-repository/golang-arch-sample/internal/foo/repository"
)

type SaveNewFooUsecase struct {
	repo repository.FooRepository
}

func (uc *SaveNewFooUsecase) Execute(input string) error {
	etty := &entity.Foo{Foo: input}
	return uc.repo.Add(etty)
}

// NewSaveNewFooUsecase contrutor do usecase
func NewSaveNewFooUsecase(repo repository.FooRepository) *SaveNewFooUsecase {
	return &SaveNewFooUsecase{
		repo: repo,
	}
}
