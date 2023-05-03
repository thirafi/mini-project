package controller

import (
	"net/http"

	"project_structure/model"
	"project_structure/usecase"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
}

type authController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *authController {
	return &authController{
		userUsecase,
	}
}

func (a *authController) LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := a.userUsecase.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}
