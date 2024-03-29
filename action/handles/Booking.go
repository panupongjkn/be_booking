package handles

import (
	"github.com/labstack/echo"
	"github.com/panupong25509/be_booking_sign/action/repositories"
	"github.com/panupong25509/be_booking_sign/models"
)

func AddBooking(c echo.Context) error {
	newBooking, err := repositories.AddBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	booking := newBooking.(models.Booking)
	return c.JSON(200, booking.ReturnJsonID())
}

func GetBookingById(c echo.Context) error {
	days, err := repositories.GetBookingById(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, days)
}

func GetBookingDayBySign(c echo.Context) error {
	days, err := repositories.GetBookingDaysBySign(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, days)
}

func ApproveBooking(c echo.Context) error {
	message, err := repositories.ApproveBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, message)
}

func RejectBooking(c echo.Context) error {
	message, err := repositories.RejectBooking(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, message)
}

func GetBookingAdmin(c echo.Context) error {
	booking, _ := repositories.GetBookingAdmin(c)
	return c.JSON(200, booking)
}

func GetBookingUser(c echo.Context) error {
	booking, err := repositories.GetBookingUser(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, booking)
}

func GetBookingByFilter(c echo.Context) error {
	booking, err := repositories.GetBookingByFilter(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, booking)
}

func GetSummaryMonth(c echo.Context) error {
	summary, err := repositories.GetSummaryMonth(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, summary)
}
func GetSummarySign(c echo.Context) error {
	summary, err := repositories.GetSummarySign(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, summary)
}
func GetSummaryOrganization(c echo.Context) error {
	summary, err := repositories.GetSummaryOrganization(c)
	if err != nil {
		status := err.(models.Error)
		return c.JSON(status.Code, status)
	}
	return c.JSON(200, summary)
}
