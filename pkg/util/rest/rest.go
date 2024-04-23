package rest

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/thhuang/kakao/pkg/util/ctx"
)

type HealthStatus string

const (
	HealthStatusUp   HealthStatus = "up"
	HealthStatusDown HealthStatus = "down"
)

type HealthResponse struct {
	Status HealthStatus `json:"status"`
}

func Serve(ctx ctx.CTX, app *fiber.App, addr string) error {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(HealthResponse{
			Status: HealthStatusUp,
		})
	})

	if err := app.Listen(addr); err != nil {
		return errors.Wrap(err, "Serve::Listen")
	}

	return nil
}

// RegisterStringTag registers a custom validation rule for a string field based on a regular expression.
func RegisterStringTag(validate *validator.Validate, tag string, regex *regexp.Regexp) {
	if err := validate.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		result := regex.MatchString(s)
		return result
	}); err != nil {
		panic(err)
	}
}

// ParseBody parses the request body into a struct and validates it.
func ParseBody(c *fiber.Ctx, validate *validator.Validate, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return err
	}
	return validateStruct(validate, obj)
}

func validateStruct(validate *validator.Validate, obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = indirect(v)
	}

	switch v.Type().Kind() {
	case reflect.Struct:
		return validate.Struct(obj)
	}

	// Simply ignore validation for other cases.
	return nil
}

func indirect(v reflect.Value) reflect.Value {
	// Recursively digs into v if the type of v is a pointer.
	for {
		if (v.Kind() == reflect.Interface && !v.IsNil()) || v.Kind() == reflect.Ptr {
			v = v.Elem()
			continue
		}
		return v
	}
}
