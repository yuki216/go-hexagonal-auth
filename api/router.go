package api

import (
	echo "github.com/labstack/echo/v4"
	"go-hexagonal-auth/api/middleware"
	"go-hexagonal-auth/api/v1/auth"
	"go-hexagonal-auth/api/v1/media"
	"go-hexagonal-auth/api/v1/user"
	"go-hexagonal-auth/config"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, authController *auth.Controller, userController *user.Controller, mediaController *media.Controller, cfg config.Config) {
	if authController == nil || userController == nil ||mediaController == nil {
		panic("Controller parameter cannot be nil")
	}

	//authentication with Versioning endpoint
	authV1 := e.Group("api/v1/auth")
	authV1.POST("/login", authController.Login)
	authV1.POST("/register-admin", authController.RegisterAdmin)
	authV1.POST("/register-user", authController.RegisterUser)

	//user with Versioning endpoint
	userV1 := e.Group("api/v1/users")
	userV1.Use(middleware.JWTMiddleware(cfg))
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:id", userController.UpdateUser)


	//user with Versioning endpoint
	mediaV1 := e.Group("api/v1/media")
	mediaV1.Use(middleware.JWTMiddleware(cfg))
	mediaV1.POST("/upload", mediaController.MediaUpload)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
