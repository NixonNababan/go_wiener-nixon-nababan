package controller

import (
	"net/http"

	"loginnregister/models/payload"
	"loginnregister/usecase"

	"github.com/labstack/echo"
)

func GetProfilController(c echo.Context) error {
	_, id := Authorization(c)
	user, err := usecase.GetProfil(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the profil",
		"data":    user,
	})
}

func UpdateProfilController(c echo.Context) error {
	var req payload.UpdateProfil
	_, id := Authorization(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateProfil(id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update profil",
	})
}

func DeleteUserController(c echo.Context) error {
	_, id := Authorization(c)

	password := c.FormValue("password")
	err := usecase.DeleteUser(id, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
