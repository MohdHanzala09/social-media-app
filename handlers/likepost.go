package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func LikePost(c *echo.Context) error {
	userid := c.Get("userID")
	// tweetid := c.Get("tweetID")
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to like post"})
	}

	if userid == "" || id <= 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "login first"})
	}

	query := `
	INSERT INTO likes (user_id , tweet_id) VALUES (? , ?)
	`
	_, err = db.Exec(query, userid, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "please retry"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"msg": "like added"})
}
