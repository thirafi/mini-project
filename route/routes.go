package route

import (
	"net/http"
	"project_structure/constant"
	"project_structure/controller"
	"project_structure/repository/database"
	"project_structure/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {

	userRepository := database.NewUserRepository(db)
	blogRepository := database.NewBlogRepository(db)
	bookRepository := database.NewBookRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, blogRepository)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)
	blogController := controller.NewBlogController(blogUsecase, blogRepository)
	bookController := controller.NewBookController(bookUsecase, bookRepository)

	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", authController.LoginUserController)

	// user collection
	user := e.Group("/users", middleware.JWT([]byte(constant.SECRET_JWT)))
	user.GET("", userController.GetUserController)
	user.GET("/:id", userController.GetUserController)
	user.POST("", userController.CreateUserController)
	user.PUT("/:id", userController.UpdateUserController)
	user.DELETE("/:id", userController.DeleteUserController)

	// book collection
	book := e.Group("/books", middleware.JWT([]byte(constant.SECRET_JWT)))
	book.GET("", bookController.GetBookController)
	book.GET("/:id", bookController.GetBookController)
	book.POST("", bookController.CreateBookController)
	book.PUT("/:id", bookController.UpdateBookController)
	book.DELETE("/:id", bookController.DeleteBookController)

	// book collection
	blog := e.Group("/blogs", middleware.JWT([]byte(constant.SECRET_JWT)))
	blog.GET("", blogController.GetBlogController)
	blog.GET("/:id", blogController.GetBlogController)
	blog.POST("", blogController.CreateBlogController)
	blog.PUT("/:id", blogController.UpdateBlogController)
	blog.DELETE("/:id", blogController.DeleteBlogController)

}
