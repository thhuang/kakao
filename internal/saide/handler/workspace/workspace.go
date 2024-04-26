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

	fileStructure := directory{
		"/": {
			file{
				"README.md": uuid.NewString(),
			},
			directory{
				"cmd": {
					file{
						"main.go": uuid.NewString(),
					},
				},
			},
			directory{
				"pkg": {
					directory{
						"service": {
							directory{
								"mongo": {
									file{
										"mongo.go": uuid.NewString(),
									},
									file{
										"impl.go": uuid.NewString(),
									},
									file{
										"impl_test.go": uuid.NewString(),
									},
								},
							},
							directory{
								"redis": {
									file{
										"redis.go": uuid.NewString(),
									},
									file{
										"impl.go": uuid.NewString(),
									},
									file{
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
