package api

import (
	"github.com/gofiber/websocket/v2"
)

func (api *API) Consume(c *websocket.Conn) {
	topic := c.Query("topic", "")
	if topic == "" {
		c.Close()
		return
	}

	messages, err := api.service.Topic.GetAllMessages(topic)
	if err != nil {
		c.Close()
		return
	}

	go func() {
		for _, m := range messages {
			mJson := m.ToJSON()
			if err := c.WriteMessage(websocket.TextMessage, mJson); err != nil {
				continue
			}

			api.service.Topic.DeleteMessage(topic, m.ID)
		}
	}()

	go func() {
		for {
			select {
			case pub := <-api.chPublishing:
				if topic == pub.Topic {
					c.WriteMessage(websocket.TextMessage, pub.Message.ToJSON())
					api.service.Topic.DeleteMessage(topic, pub.Message.ID)
				}
			}
		}
	}()

	for {
	}

}
