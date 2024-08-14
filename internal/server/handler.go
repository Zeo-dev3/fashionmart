package server

import (
	authHttp "github.com/Zeo-dev3/fashionmart/internal/auth/delivery/http"
	authRepository "github.com/Zeo-dev3/fashionmart/internal/auth/repository"
	authUsecase "github.com/Zeo-dev3/fashionmart/internal/auth/usecase"
	"github.com/Zeo-dev3/fashionmart/internal/middleware"
	productHttp "github.com/Zeo-dev3/fashionmart/internal/product/delivery/http"
	productRepository "github.com/Zeo-dev3/fashionmart/internal/product/repository"
	productUsecase "github.com/Zeo-dev3/fashionmart/internal/product/usecase"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlres(app *fiber.App) error {
	// setup repository
	authRepo := authRepository.NewAuthRepo(s.db)
	productRepo := productRepository.NewProductRepository(s.db)

	// setup usecase
	authUC := authUsecase.NewAuthUseCase(authRepo, s.cfg)
	productUC := productUsecase.NewProductUsecas(productRepo)

	// setup handler
	authHandler := authHttp.NewAuthHandler(authUC)
	productHandler := productHttp.NewProductHandler(productUC)

	// setup middleware
	authMiddleware := middleware.AuthJwtMiddleware(s.cfg)
	adminMiddleware := middleware.AdminRoleAuth()

	// swagger route
	swaggerCfg := swagger.Config{
		BasePath: s.cfg.Swagger.BasePath,
		FilePath: s.cfg.Swagger.FilePath,
		Path: s.cfg.Swagger.Path,
		Title: s.cfg.Swagger.Title,		
	}
	s.app.Use(swagger.New(swaggerCfg))

	authHttp.MapAuthRoutes(s.app, authHandler, authMiddleware, adminMiddleware)
	productHttp.MapProductRoutes(s.app, productHandler, authMiddleware, adminMiddleware)

	return nil
}
