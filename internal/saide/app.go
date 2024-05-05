package saide

import (
	"flag"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/thhuang/kakao/pkg/util/ctx"
	"github.com/thhuang/kakao/pkg/util/rest"

	// Services
	workspaceService "github.com/thhuang/kakao/internal/saide/service/workspace"

	// Handlers
	workspaceHandler "github.com/thhuang/kakao/internal/saide/handler/workspace"
)

var (
	port = flag.Int("port", 80, "api server port")

	limiterConfig = limiter.Config{Max: 10, Expiration: 30 * time.Second}
)

type App struct {
	app *fiber.App
}

func New(ctx ctx.CTX) *App {
	app := fiber.New()

	// Default middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiterConfig))
	app.Use(rest.AddContext())

	// Services
	workspaceService := workspaceService.New()

	// Handlers
	workspaceHandler.New(app.Group("/workspaces"), workspaceService)

	return &App{
		app: app,
	}
}

func (a *App) Run(ctx ctx.CTX) {
	ctx.Infof("Start serving on port %d", *port)
	if err := rest.Serve(ctx, a.app, fmt.Sprintf(":%d", *port)); err != nil {
		ctx.WithError(err).Panicf("rest.Serve failed")
	}
}
