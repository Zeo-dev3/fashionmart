package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Zeo-dev3/fashionmart/config"
	"github.com/Zeo-dev3/fashionmart/internal/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{
		AppName: cfg.Server.AppName,
		Prefork: cfg.Server.Prefork,
	})
	return &Server{app: app, cfg: cfg, db: db}
}

func (s *Server) Run() error {
	if err := s.MapHandlres(s.app); err != nil {
		return err
	}

	s.db.Exec("CREATE EXTENSION IF NOT EXISTS\"uuid-ossp\"")

	if err := entity.MigrateUserEntity(s.db); err != nil {
		log.Println("failed to migrate user entity", err)
	}

	go func() {
		s.app.Listen(s.cfg.Server.Port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	<-ctx.Done()
	return s.app.Shutdown()

}
