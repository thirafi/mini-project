package usecase

import (
	"fmt"
	"project_structure/middleware"
	"project_structure/model"
	"project_structure/repository/database"
)

type UserUsecase interface {
	LoginUser(user *model.User) (err error)
	CreateUser(user *model.User) error
	GetUser(id uint) (user model.User, err error)
	GetListUsers() (users []model.User, err error)
	UpdateUser(user *model.User) (err error)
	DeleteUser(id uint) (err error)
}

type userUsecase struct {
	userRepository database.UserRepository
	blogRepository database.BlogRepository
}

func NewUserUsecase(
	userRepo database.UserRepository,
	blogRepo database.BlogRepository,
) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
		blogRepository: blogRepo,
	}
}

func (u *userUsecase) LoginUser(user *model.User) (err error) {
	// check to db email and password
	err = u.userRepository.LoginUser(user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	// generate jwt
	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println("GetUser: Error Generate token")
		return
	}
	user.Token = token
	return
}

func (u *userUsecase) CreateUser(user *model.User) error {

	err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetUser(id uint) (user model.User, err error) {
	user.ID = id
	err = u.userRepository.GetUser(&user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	blog, err := u.blogRepository.GetBlogsByUserId(id)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	user.Blogs = append(user.Blogs, blog)
	return
}

func (u *userUsecase) GetListUsers() (users []model.User, err error) {
	users, err = u.userRepository.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers: Error getting users from database")
		return
	}
	return
}

func (u *userUsecase) UpdateUser(user *model.User) (err error) {
	err = u.userRepository.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func (u *userUsecase) DeleteUser(id uint) (err error) {
	user := model.User{}
	user.ID = id
	err = u.userRepository.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
