package http

import (
	"gogrpc/pkg/blog"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) createArticle(c *fiber.Ctx) error {
	a := new(blog.Article)
	err := c.BodyParser(a)
	if err != nil {
		return err
	}

	errCreate := s.articleService.Create(a)
	if errCreate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errCreate.Error())
	}

	return c.JSON(a)
}

func (s *Server) updateArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	a, errGet := s.articleService.Get(id) // TODO handle partial update also
	if errGet != nil {
		return errGet
	}

	errBody := c.BodyParser(a)
	if errBody != nil {
		return errBody
	}

	errUpdate := s.articleService.Update(a)
	if errUpdate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errUpdate.Error())
	}

	return c.JSON(a)
}

func (s *Server) deleteArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	errDelete := s.articleService.Delete(id)
	if errDelete != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errDelete.Error())
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) getArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	a, errGet := s.articleService.Get(id)
	if errGet != nil {
		return c.Status(fiber.StatusNotFound).JSON(errGet.Error())
	}

	return c.JSON(a)
}
