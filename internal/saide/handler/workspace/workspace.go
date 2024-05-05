package workspace

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/thhuang/kakao/internal/saide/service/workspace"
	"github.com/thhuang/kakao/pkg/util/ctx"
	"github.com/thhuang/kakao/pkg/util/errutil"
	"github.com/thhuang/kakao/pkg/util/keyword"
)

var (
	ErrInvalidFileStructure = errors.New("invalid file structure")
)

func New(
	r fiber.Router,
	workspaceService workspace.Service,
) {
	h := handler{
		workspaceService: workspaceService,
	}

	r.Get(":id", h.getWorkspace)
}

type handler struct {
	workspaceService workspace.Service
}

func (h *handler) getWorkspace(c *fiber.Ctx) error {
	context := c.Locals(keyword.CTX).(ctx.CTX)

	workspaceIdString := c.Params("id")
	workspaceId, err := uuid.Parse(workspaceIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errutil.ErrorResponse{
			Code:    errutil.ErrorCodeBadRequest,
			Message: err.Error(),
		})
	}

	workspace, err := h.workspaceService.GetWorkspace(context, workspaceId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errutil.ErrorResponse{
			Code:    errutil.ErrorCodeUnknown,
			Message: err.Error(),
		})
	}

	res, err := NewGetWorkspaceResponse(workspace)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(errutil.ErrorResponse{
			Code:    errutil.ErrorCodeUnknown,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
