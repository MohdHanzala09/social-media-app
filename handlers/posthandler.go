package handlers

import (
	// "database/sql"
	"net/http"

	"github.com/MohdHanzala09/social-media-app/models"
	"github.com/labstack/echo/v5"
)

func PostTweet(c *echo.Context) error {
	var tweet models.Tweets
	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid body"})
	}
	query := `
	INSERT INTO likes (user_id, content) VALUES (?,?)
	`
	tweet.UserId = c.Get("UserID").(int)

	_, err := db.Exec(query, tweet.UserId, tweet.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to post tweet"})
	}
	
	return c.JSON(http.StatusCreated, map[string]string{"msg":"created"})
}
