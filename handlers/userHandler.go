package handlers

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/MohdHanzala09/social-media-app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetUser(c *echo.Context) error {
	return nil
}

func GetUserByID(c *echo.Context) error {
	return nil
}

func LoginUser(c *echo.Context) error {
	// binding the struct
	var user models.LoginSchema
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	// quering database
	var userId int
	var userPass string
	query := `SELECT id ,password FROM users WHERE email=?`
	err := db.QueryRow(query, user.Email).Scan(&userId, &userPass)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "server error"})
	}

	//comparing password
	err = bcrypt.CompareHashAndPassword([]byte(userPass) , []byte(user.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error":"wrong email or password"})
	}

	//signing the token
	claims := &models.Claims{
		Id:    userId,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	signtoken , err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError , map[string]string{"error":"could not create token"})
	}

	//seting the cookie
	c.SetCookie(&http.Cookie{
		Name: "access-token",
		Value: signtoken,
		HttpOnly: true,
		Secure: true,
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Expires: time.Now().Add(30 * time.Minute),
	})

	return c.JSON(http.StatusOK, map[string]string{"msg": "login successful"})
}

func CreateUser(c *echo.Context) error {
	var user models.UserSchema
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid body"})
	}

	if user.Name == "" || user.Email == ""  || user.Password == ""{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid body"})
	}

	if len(user.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"password should atleast 8 char or more"})
	}
	hashedPass , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":"server error"})
	}

	query := `INSERT INTO users (name,email,password,dob) VALUES (?,?,?,?)`

	res, err := db.Exec(query, user.Name, user.Email, string(hashedPass) ,user.DOB)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "user creation failed"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "user creation failed"})
	}

	user.ID = int(id)

	return c.JSON(http.StatusCreated, user)
}
