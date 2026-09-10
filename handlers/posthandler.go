package handlers

import (
	// "database/sql"
	"net/http"

	"github.com/MohdHanzala09/social-media-app/models"
	"github.com/labstack/echo/v5"
)

// var dbPost *sql.DB
// func SetDBPost(database *sql.DB){
// 	dbPOst = database
// }

func PostTweet(c *echo.Context) error {
	var tweet models.Posts
	if err := c.Bind(&tweet); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid body"})
	}
	query := `
	INSERT INTO likes (tweet,userid,useremail) VALUES (?,?,?)
	`
	tweet.UserID = c.Get("UserID").(int)
	tweet.UserEmail = c.Get("UserEmail").(string)

	_, err := db.Exec(query, tweet.Tweet, tweet.UserID, tweet.UserEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to post tweet"})
	}
	
	return c.JSON(http.StatusCreated, map[string]string{"msg":"created"})
}
