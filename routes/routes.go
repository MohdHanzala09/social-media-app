package routes

import (
	// "log"

	"net/http"

	"github.com/MohdHanzala09/social-media-app/handlers"
	"github.com/MohdHanzala09/social-media-app/middleware"
	"github.com/labstack/echo/v5"
)

func RegisterAllUserRoutes(e *echo.Echo) {
	
	//users Route
	e.POST("/users" , handlers.CreateUser)
	e.GET("/users" , handlers.GetUser) //admin use
	e.GET("/users/:id" , handlers.GetUserByID)

	e.GET("/" , func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"msg":"hello world"})
	})

}

func RegisterAllPostsRoutes(e *echo.Echo) {
	e.POST("/users/posts" , handlers.PostTweet , middleware.CheckToken)
}
