package main

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/boil"
	"go_app/api/routes"
	"go_app/pkg/todo"
	"log"
	"os"
	"time"
)

func connectDB() {
	err := godotenv.Load("../conf/.env")
	if err != nil {
		panic(err)
	}

	jst, _ := time.LoadLocation(os.Getenv("LOC"))
	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("USER"),
		Passwd:    os.Getenv("PASS"),
		Addr:      os.Getenv("ADDR"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_0900_ai_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	boil.SetDB(db)
	boil.DebugMode = true
}

func main() {
	// DBコネクションを取得する
	connectDB()

	app := fiber.New()
	validate := validator.New()
	app.Use(cors.New())

	// routing
	api := app.Group("/api")
	routes.TodoRouter(api, todo.NewService(todo.NewRepo()), validate)

	log.Fatal(app.Listen(":8080"))
}
