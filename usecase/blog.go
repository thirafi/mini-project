package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type BlogUsecase interface {
	CreateBlog(blog *model.Blog) error
	GetBlog(id uint) (blog model.Blog, err error)
	GetListBlogs() (blogs []model.Blog, err error)
	UpdateBlog(blog *model.Blog) (err error)
	DeleteBlog(id uint) (err error)
}

type blogUsecase struct {
	blogRepository database.BlogRepository
}

func NewBlogUsecase(blogRepo database.BlogRepository) *blogUsecase {
	return &blogUsecase{blogRepository: blogRepo}
}

func (b *blogUsecase) CreateBlog(blog *model.Blog) error {

	// check name cannot be empty
	if blog.Title == "" {
		return errors.New("blog title cannot be empty")
	}

	// check email
	if blog.Content == "" {
		return errors.New("blog content cannot be empty")
	}

	err := b.blogRepository.CreateBlog(blog)
	if err != nil {
		return err
	}

	return nil
}

func (b *blogUsecase) GetBlog(id uint) (blog model.Blog, err error) {
	blog, err = b.blogRepository.GetBlog(id)
	if err != nil {
		fmt.Println("GetBlog: Error getting blog from database")
		return
	}
	return
}

func (b *blogUsecase) GetListBlogs() (blogs []model.Blog, err error) {
	blogs, err = b.blogRepository.GetBlogs()
	if err != nil {
		fmt.Println("GetListBlogs: Error getting blogs from database")
		return
	}
	return
}

func (b *blogUsecase) UpdateBlog(blog *model.Blog) (err error) {
	err = b.blogRepository.UpdateBlog(blog)
	if err != nil {
		fmt.Println("UpdateBlog : Error updating blog, err: ", err)
		return
	}

	return
}

func (b *blogUsecase) DeleteBlog(id uint) (err error) {
	blog := model.Blog{}
	blog.ID = id
	err = b.blogRepository.DeleteBlog(&blog)
	if err != nil {
		fmt.Println("DeleteBlog : error deleting blog, err: ", err)
		return
	}

	return
}
