package api

import (
	"github.com/gofiber/fiber/v2"
)

type PublishMessageReq struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}

func (api *API) Publish(c *fiber.Ctx) error {

	input := PublishMessageReq{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}

	message, err := api.service.Topic.Publish(input.Topic, input.Message)
	if err != nil {
		return err
	}

	api.chPublishing <- publishing{
		Topic:   input.Topic,
		Message: message,
	}

	return c.JSON(fiber.Map{
		"message": "sent",
		"topic":   input.Topic,
		"id":      message.ID,
	})
}
