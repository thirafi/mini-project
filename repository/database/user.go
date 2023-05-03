package database

import (
	"project_structure/config"
	"project_structure/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUsers() (users []model.User, err error)
	GetUser(user *model.User) (err error)
	GetUserWithBlog(id uint) (user model.User, err error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
	LoginUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUsers() (users []model.User, err error) {
	if err = config.DB.Model(&model.User{}).Preload("Blogs").Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) GetUser(user *model.User) (err error) {
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) GetUserWithBlog(id uint) (user model.User, err error) {
	user.ID = id
	if err = config.DB.Model(&model.User{}).Preload("Blogs").First(&user).Error; err != nil {
		return
	}
	return
}

func (u *userRepository) UpdateUser(user *model.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUser(user *model.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) LoginUser(user *model.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
