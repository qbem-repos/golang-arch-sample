package usecase

import (
	"errors"
	"testing"

	"github.com/qbem-repository/golang-arch-sample/internal/foo/entity"
)

type MockFooRepository struct{}

func (m *MockFooRepository) Add(etty *entity.Foo) error {
	return errors.New("error")
}

func TestSaveNewFooUsecase_WhenInputIsNull_ThenReturnError(t *testing.T) {
	r := new(MockFooRepository)
	s := NewSaveNewFooUsecase(r)

	if err := s.Execute(""); err == nil {
		t.Errorf("Error must not be nil")
	}
}
