package server

import (
	"github.com/Zeo-dev3/fashionmart/internal/auth/delivery/http"
	"github.com/Zeo-dev3/fashionmart/internal/auth/repository"
	"github.com/Zeo-dev3/fashionmart/internal/auth/usecase"
	"github.com/Zeo-dev3/fashionmart/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlres(app *fiber.App) error {
	authRepo := repository.NewAuthRepo(s.db)

	authUC := usecase.NewAuthUseCase(authRepo, s.cfg)

	authHandler := http.NewAuthHandler(authUC)

	authMiddleware := middleware.AuthJwtMiddleware(s.cfg)
	adminMiddleware := middleware.AdminRoleAuth()

	http.NewAuthRoute(s.app, authHandler, authMiddleware,adminMiddleware)

	return nil
}
