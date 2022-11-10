package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null/v8"
	_ "github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go_app/pkg/models"
	"log"
	"time"
)

func main() {
	//http.HandleFunc("/", echoHello)
	//http.ListenAndServe(":8000", nil)
	app := fiber.New()

	app.Get("", func(c *fiber.Ctx) error {
		con := connectDB()
		ctx := context.Background()

		//result, _ := models.FindTodo(context.Background(), con, 1)
		//result, _ := models.Todos().All(ctx, con)

		type Custom struct {
			ID   int         `boil:"id" json:"id" toml:"id" yaml:"id"`
			Name null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
		}
		var result []Custom
		err := models.Todos(
			qm.Select("todos.id, users.name"),
			qm.LeftOuterJoin("users on todos.userId = users.id")).Bind(ctx, con, &result)
		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(result)
	})

	log.Fatal(app.Listen(":8080"))
}

func connectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	c := mysql.Config{
		DBName:    "go_app_test",
		User:      "go_test",
		Passwd:    "pass",
		Addr:      "db",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}
