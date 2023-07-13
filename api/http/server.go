package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Server struct {
	db  *gorm.DB
	app *fiber.App
}

func NewServer(db *gorm.DB) *Server {
	app := fiber.New()
	server := &Server{
		db:  db,
		app: app,
	}

	server.setupMiddlewares()
	server.setupRoutes()

	return server
}

func (s *Server) Run() error {
	return s.app.Listen(":3000") // TODO read the port from the env
}

func (s *Server) setupRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func (s *Server) setupMiddlewares() {
	s.app.Use(logger.New()) // TODO force the logger middleware to use the zerolog
}
