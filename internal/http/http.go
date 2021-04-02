package http

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/dentych/dinner-dash/internal/api"
	"gitlab.com/dentych/dinner-dash/internal/http/internal/handler"
)

type Server struct {
	userApi       *api.UserApi
	familyApi     *api.FamilyApi
	ingredientApi *api.IngredientApi
	recipeApi     *api.RecipeApi
}

func NewServer(userApi *api.UserApi, familyApi *api.FamilyApi) *Server { //, familyApi api.FamilyApi, ingredientApi api.IngredientApi, recipeApi api.RecipeApi) *server {
	return &Server{
		userApi: userApi,
		familyApi: familyApi,
		//ingredientApi: ingredientApi,
		//recipeApi: recipeApi,
	}
}

func (h *Server) SetupAndStart(ctx context.Context, jwtCertificate string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	jwtMiddleware := setupJwt(e, jwtCertificate)
	apiGroup := e.Group("/api", jwtMiddleware)
	apiGroup.GET("/user", handler.GetUser(h.userApi))
	apiGroup.POST("/user", handler.CreateUser(h.userApi))
	apiGroup.GET("/family/:id", handler.GetFamily(h.familyApi))
	apiGroup.POST("/family", handler.CreateFamily(h.familyApi))
	apiGroup.PUT("/family/:id", handler.UpdateFamily(h.familyApi))
	apiGroup.POST("/family/:id/invite", handler.GenerateInvitation(h.familyApi))
	apiGroup.PUT("/invite/:invitationId", handler.AcceptInvitation(h.familyApi))
	apiGroup.GET("/recipes", handler.GetRecipes)

	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func setupJwt(e *echo.Echo, jwtCertificate string) echo.MiddlewareFunc {
	pemCert, _ := pem.Decode([]byte(jwtCertificate))
	if pemCert == nil {
		e.Logger.Fatal("Failed to parse pem certificate")
		return nil
	}
	certificate, err := x509.ParseCertificate(pemCert.Bytes)
	if err != nil {
		e.Logger.Fatal("Failed to parse pemCert.Bytes", err)
	}
	signingKey := certificate.PublicKey
	config := middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodRS256.Alg(),
		SigningKey:    signingKey,
	}
	return middleware.JWTWithConfig(config)
}
