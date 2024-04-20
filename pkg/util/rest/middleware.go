package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/thhuang/kakao/util/ctx"
	"github.com/thhuang/kakao/util/keyword"
)

func AddContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Add a context with requestId.
		context := ctx.WithValue(ctx.Background(), keyword.RequestId, uuid.New())
		c.Locals(keyword.CTX, context)

		// Continue the stack.
		return c.Next()
	}
}
