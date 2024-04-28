package workspace

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	errorUtils "github.com/thhuang/kakao/pkg/util/error"
)

func New(
	r fiber.Router,
) {
	h := handler{}

	r.Get(":id", h.getWorkspace)
}

type handler struct {
}

func (h *handler) getWorkspace(c *fiber.Ctx) error {
	// context := c.Locals(keyword.CTX).(ctx.CTX)

	workspaceIdString := c.Params("id")
	workspaceId, err := uuid.Parse(workspaceIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorUtils.ErrorResponse{
			Code:    errorUtils.ErrorCodeBadRequest,
			Message: err.Error(),
		})
	}

	fileStructure := Directory{
		"/": {
			File{
				"README.md": uuid.NewString(),
			},
			Directory{
				"cmd": {
					File{
						"main.go": uuid.NewString(),
					},
				},
			},
			Directory{
				"pkg": {
					Directory{
						"service": {
							Directory{
								"mongo": {
									File{
										"mongo.go": uuid.NewString(),
									},
									File{
										"impl.go": uuid.NewString(),
									},
									File{
										"impl_test.go": uuid.NewString(),
									},
								},
							},
							Directory{
								"redis": {
									File{
										"redis.go": uuid.NewString(),
									},
									File{
										"impl.go": uuid.NewString(),
									},
									File{
										"impl_test.go": uuid.NewString(),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !fileStructure.isValid() {
		return c.Status(fiber.StatusInternalServerError).JSON(errorUtils.ErrorResponse{
			Code:    errorUtils.ErrorResponseValidationFailed,
			Message: "invalid file structure",
		})
	}

	return c.Status(fiber.StatusOK).JSON(getWorkspaceResponse{
		Id:            workspaceId.String(),
		FileStructure: fileStructure,
	})
}
