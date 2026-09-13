package handlers

import (
	"net/http"
	"strconv"

	"github.com/MohdHanzala09/social-media-app/models"
	"github.com/labstack/echo/v5"
)

func Comment(c *echo.Context) error {
	var comments models.Comments
	userid := c.Get("userID")
	tweetid , err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error":"failed"})
	}
	comments.TweetId = tweetid
	comments.UserId = userid.(int)
	err = c.Bind(&comments)
	if err != nil {
		c.JSON(http.StatusInternalServerError , map[string]string{"error":"failed to comment"})
	}
	
	
	query := `
	INSERT INTO comments (tweet_id , user_id, content) VALUES (? ,? ,?)
	`
	_ , err = db.Exec(query , comments.TweetId , comments.UserId , comments.Content)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"failed to comment"})
	}

	return c.JSON(http.StatusCreated , map[string]string{"msg":"comment posted"})
}