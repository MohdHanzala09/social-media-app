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

	if tweet.Content == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "content is required"})
	}

	query := `
	INSERT INTO tweets (user_id, content) VALUES (?,?)
	`
	tweet.UserId = c.Get("UserID").(int)

	res, err := db.Exec(query, tweet.UserId, tweet.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to post tweet"})
	}

	tweetID ,err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError , map[string]string{"error":"plz repost"})
	}
	
	c.Set("tweetID" , tweetID)
	return c.JSON(http.StatusCreated, map[string]string{"msg": "created"})
}
