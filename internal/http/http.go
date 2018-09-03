package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/dentych/dinner-dash/internal/http/internal/handler"
)

func Setup(ctx context.Context) {
	e := echo.New()

	jwtConfig := middleware.JWTConfig{
		SigningKey: "9O8gk65hm61VZlpRUlqtLliGDFovyG6m",
	}
	apiGroup := e.Group("/api", middleware.JWTWithConfig(jwtConfig))
	apiGroup.GET("/recipes", handler.GetRecipes)

	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
