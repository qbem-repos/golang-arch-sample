package boundary

import (
	"io"

	"github.com/qbem-repository/golang-arch-sample/internal/foo/entity"
)

type CreateFooInput struct {
	Message string
}

func (i *CreateFooInput) FromJson(r *io.Reader) {
    
}

func (i *CreateFooInput) Entity() *entity.Foo {
	return &entity.Foo{
		Foo: i.Message,
	}
}

type CreateFooOutput struct{}
