package controller

import (
	"net/http"
	"project_structure/model"
	"project_structure/model/payload"
	"project_structure/repository/database"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetUsercontroller(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
}

type userController struct {
	userUsecase    usecase.UserUsecase
	userRepository database.UserRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
) *userController {
	return &userController{
		userUsecase,
		userRepository,
	}
}

func (u *userController) GetUsercontroller(c echo.Context) error {
	users, e := u.userRepository.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (u *userController) GetUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := u.userRepository.GetUserWithBlog(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// create new user
func (u *userController) CreateUserController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if err := c.Validate(payload); err != nil {
		return err
	}
	user := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
	if err := u.userUsecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func (u *userController) DeleteUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := u.userUsecase.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func (u *userController) UpdateUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := model.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := u.userUsecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
