package workspace

import (
	"github.com/google/uuid"
	"github.com/thhuang/kakao/pkg/util/ctx"
)

type Service interface {
	GetWorkspace(context ctx.CTX, id uuid.UUID) (*Workspace, error)
}
