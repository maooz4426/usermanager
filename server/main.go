package main

import (
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/usermanager/infra/wire"
	"github.com/maooz4426/usermanager/lib/usermanage"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("can not read env: %v", err)
	}

	mysqlUser := os.Getenv("MYSQL_USER")         // "user"
	mysqlPassword := os.Getenv("MYSQL_PASSWORD") // "password"
	mysqlDatabase := os.Getenv("MYSQL_DATABASE") // "db"

	dsn := fmt.Sprintf("%s:%s@tcp(usermanage_db:3306)/%s?parseTime=true",
		mysqlUser,
		mysqlPassword,
		mysqlDatabase,
	)

	e := echo.New()
	conn, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		panic(err)
	}

	sess := conn.NewSession(nil)

	handler := wire.Wire(sess)

	usermanage.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":8080"))
}
