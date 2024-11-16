package manager

import (
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
	"test/project/internal/posts"
	"test/project/pkg/config"
)

const (
	baseURL        = "/api/test/v1"
	healthCheckURL = "/health"
	videosURL      = "/videos"
)

func Manager(db *gocql.Session, cfg *config.Configs) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	router := app.Group(baseURL)

	router.Get(healthCheckURL, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "The Service is Running, YEAH!",
		})
	})

	videoRouterManager := router.Group(videosURL)
	videoRouterRepository := posts.NewRepository(db)
	videoRouterHandler := posts.NewHandler(videoRouterRepository, cfg)
	videoRouterHandler.Register(videoRouterManager)

	return app
}
