package handles

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/panupong25509/be_booking_sign/action/repositories"
	"github.com/panupong25509/be_booking_sign/models"
)

// func Register(c echo.Context) error {
// 	data := DynamicPostForm(c)
// 	success, err := repositories.Register(c, data)
// 	if err != nil {
// 		status := err.(models.Error)
// 		// return c.Render(status.Code, r.JSON(status))
// 	}
// 	// return c.Render(200, r.JSON(success))
// }

func Login(c echo.Context) error {
	// data := DynamicPostForm(c)
	// _, _ := repositories.Login(c, data)
	log.Print("login")
	return c.JSON(http.StatusOK, "users")
}

func Register(c echo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.Register(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, success)
}

func GetUserByUsername(c echo.Context) error {
	data := DynamicPostForm(c)
	success, err := repositories.GetUserByUsername(c, data)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, success)
}

func GetUserById(c echo.Context) error {
	success, err := repositories.GetUserById(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, success)
}
