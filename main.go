package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

const role = "admin"

func main() {
	s := echo.New()

	s.Use(mw)

	s.GET("/status", handel)

	if err := s.StartServer(&http.Server{Addr: ":8080"}); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}

func handel(c echo.Context) error {
	data := time.Date(2025, time.January, 0, 0, 0, 0, 0, time.UTC)

	dur := int64(time.Until(data).Hours() / 24)

	s := fmt.Sprintf("уже прошло %d дней", dur)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func mw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		userRole := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(userRole, role) {

			fmt.Println("red button user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
