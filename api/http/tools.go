package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getIDFromURL(c *fiber.Ctx) (id uuid.UUID, err error) {
	params := struct {
		id uuid.UUID `params:"id"`
	}{}

	err = c.ParamsParser(&params)
	if err != nil {
		return
	}

	id = params.id
	return
}
