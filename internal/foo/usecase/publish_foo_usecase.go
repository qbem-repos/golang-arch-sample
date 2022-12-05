package usecase

import (
	"github.com/qbem-repository/golang-arch-sample/internal/foo/publisher"
)

func PublishNewFooUsecase(input CreateFooInbound, repo publisher.FooPublisher) (CreateFooOutbound, error) {
	repo.Add(input.Entity())
}
