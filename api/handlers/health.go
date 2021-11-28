package handlers

import "github.com/labstack/echo/v4"

func HealthCheck(c echo.Context) error {
	res := map[string]string{"Status": "UP"}
	return c.JSON(200, res)
}
