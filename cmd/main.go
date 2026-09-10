package main

import (
	"fmt"
	"log"
	"os"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/MohdHanzala09/social-media-app/db"
	"github.com/MohdHanzala09/social-media-app/handlers"
	"github.com/MohdHanzala09/social-media-app/routes"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error connecting godorenv : %v" , err)
	}
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	dsn := os.Getenv("DB_DSN")
	fmt.Println(dsn)
	db := db.ConnectDB(dsn)
	defer db.Close()
	handlers.SetDB(db)

	routes.RegisterAllUserRoutes(e)
	routes.RegisterAllPostsRoutes(e)


	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Error while starting Server : %v" , err)
	}

}
