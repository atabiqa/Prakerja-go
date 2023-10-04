package user

import (
	"net/http"
	"prakerja11/configs"
	basemodel "prakerja11/models/base"
	usermodel "prakerja11/models/user"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []usermodel.User
	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Succes",
		Data:    users,
	})
}
