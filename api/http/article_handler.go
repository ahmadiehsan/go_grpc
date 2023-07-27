package http

import (
	"gogrpc/internal/blog"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) createArticle(c *fiber.Ctx) error {
	a := new(blog.Article)
	err := c.BodyParser(a)
	if err != nil {
		return err
	}

	err = s.articleService.Create(a)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(a)
}

func (s *Server) updateArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	a, err := s.articleService.Get(id) // TODO handle partial update also
	if err != nil {
		return err
	}

	err = c.BodyParser(a)
	if err != nil {
		return err
	}

	err = s.articleService.Update(a)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(a)
}

func (s *Server) deleteArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	err = s.articleService.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) getArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	a, err := s.articleService.Get(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.JSON(a)
}
