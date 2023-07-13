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

	s.articleService.Create(a)

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

	s.articleService.Update(a)

	return c.JSON(a)
}

func (s *Server) deleteArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	s.articleService.Delete(id)

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) getArticle(c *fiber.Ctx) error {
	id, err := getIDFromURL(c)
	if err != nil {
		return err
	}

	a, errGet := s.articleService.Get(id)
	if errGet != nil {
		return errGet
	}

	return c.JSON(a)
}
