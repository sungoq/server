package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

type API struct {
	app *fiber.App

	addr string
}

func New(options ...OptionFunc) *API {
	api := &API{
		app: fiber.New(),
	}

	for _, opt := range options {
		opt(api)
	}

	api.app.Use(
		cache.New(
			cache.Config{
				Next: func(c *fiber.Ctx) bool {
					return strings.Contains(c.Route().Path, "/consume")
				},
			},
		),
	)

	if api.addr == "" {
		api.addr = ":8080"
	}

	return api
}

func (api *API) Route() {
	api.app.Post("/topics", api.CreateTopics)
	api.app.Get("/topics", api.GetAllTopics)
	api.app.Get("/topics/:id", api.GetTopicsByID)
	api.app.Delete("/topics/:id", api.DeleteTopics)

	api.app.Post("/publish", api.Publish)
	api.app.Get("/consume", api.Consume)
}

func (api *API) Start() {
	api.app.Listen(api.addr)
}
