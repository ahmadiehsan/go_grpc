package http

import (
	"gogrpc/internal/blog"
	"gogrpc/internal/conf"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	settings       conf.Settings
	app            *fiber.App
	articleService *blog.ArticleService
}

func NewServer(settings conf.Settings, as *blog.ArticleService) *Server {
	app := fiber.New()
	server := &Server{
		settings:       settings,
		app:            app,
		articleService: as,
	}

	server.setupMiddlewares()
	server.setupRoutes()

	return server
}

func (s *Server) Run() error {
	return s.app.Listen(s.settings.HTTPPort)
}

func (s *Server) setupRoutes() {
	api := s.app.Group("/api")

	articles := api.Group("/v1/articles")
	articles.Post("", s.createArticle)
	articles.Put("/:id", s.updateArticle)
	articles.Get("/:id", s.getArticle)
	articles.Delete("/:id", s.deleteArticle)
}

func (s *Server) setupMiddlewares() {
	s.app.Use(logger.New()) // TODO force the logger middleware to use the zerolog
}
