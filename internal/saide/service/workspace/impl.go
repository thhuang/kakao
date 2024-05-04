package workspace

import (
	"github.com/google/uuid"

	"github.com/thhuang/kakao/pkg/util/ctx"
)

type impl struct{}

func New() Service {
	return &impl{}
}

func (im *impl) GetWorkspace(context ctx.CTX, id uuid.UUID) (*Workspace, error) {
	// TODO: implement

	return &Workspace{}, nil
}
