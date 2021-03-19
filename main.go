package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		req := c.Request()
		ip := req.Header.Get("X-Real-Ip")
		if ip == ""{
			ip = req.Header.Get("X-Forwarded-For")
			if ip == ""{
				ip = strings.Split(req.RemoteAddr, ":")[0]
			}
		}
		return c.String(http.StatusOK, ip)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
