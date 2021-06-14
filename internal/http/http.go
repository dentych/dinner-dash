package http

import (
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

func (s *Server) SetupAndStart() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	authGroup := e.Group("/api/auth")
	authGroup.POST("/login", handler.Login(s.userApi))
	authGroup.POST("/register", handler.Register(s.userApi))
	jwtMiddleware := middleware.JWT([]byte("supersecretkey"))
	apiGroup := e.Group("/api", jwtMiddleware)
	apiGroup.GET("/user", handler.GetUser(s.userApi))
	apiGroup.GET("/family/:id", handler.GetFamily(s.familyApi))
	apiGroup.POST("/family", handler.CreateFamily(s.familyApi))
	apiGroup.PUT("/family/:id", handler.UpdateFamily(s.familyApi))
	apiGroup.POST("/family/:id/invite", handler.GenerateInvitation(s.familyApi))
	apiGroup.DELETE("/family/:id/invite", handler.DeleteInvitationLink(s.familyApi))
	apiGroup.PUT("/family/:id/leave", handler.LeaveFamily(s.familyApi))
	apiGroup.GET("/invite/:invitationId", handler.GetInvitationInformation(s.familyApi))
	apiGroup.PUT("/invite/:invitationId", handler.AcceptInvitation(s.familyApi))
	apiGroup.GET("/recipes", handler.GetRecipes)
	apiGroup.POST("recipes", handler.AddRecipe(s.recipeApi))

	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
