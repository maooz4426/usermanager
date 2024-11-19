package main

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/usermanager/infra/wire"
	"github.com/maooz4426/usermanager/lib/usermanage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()
	conn, err := dbr.Open("mysql", "user:password@tcp(usermanage_db:3306)/db", nil)
	if err != nil {
		panic(err)
	}

	sess := conn.NewSession(nil)

	handler := wire.Wire(sess)

	usermanage.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":8080"))
}
