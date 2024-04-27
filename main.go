package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome  world")
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "message")
}
func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "message")
}
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "message")
}
func main() {

	e := echo.New()
	e.GET("/", Home)
	e.POST("/register", saveUser)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
