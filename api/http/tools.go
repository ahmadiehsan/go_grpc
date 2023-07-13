package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getIDFromURL(c *fiber.Ctx) (uuid.UUID, error) {
	params := struct {
		id uuid.UUID `params:"id"`
	}{}
	err := c.ParamsParser(&params)
	if err != nil {
		return uuid.Nil, err
	}

	return params.id, nil
}
