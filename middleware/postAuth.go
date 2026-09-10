package middleware

import (
	"net/http"
	"os"

	"github.com/MohdHanzala09/social-media-app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		cookie, err := c.Cookie("access-token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{"error": "login first"})
		}
		//raw string parse at login
		tokenstr := cookie.Value

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expire token"})
		}

		c.Set("UserID", claims.Id)
		c.Set("UserEmail", claims.Email)

		return next(c)
	}
}
