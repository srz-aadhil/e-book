package app

import (
	"ebookmod/controller"
	"ebookmod/repo"
	"ebookmod/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func APIRouter(db *gorm.DB) chi.Router {

	//Author
	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	//Blog
	bookRepo := repo.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	//User
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()
	r.Route("/books", func(r chi.Router) {
		r.Post("/create", bookController.CreateBook)
		r.Get("/{id}", bookController.GetBook)
		r.Get("/", bookController.GetAllBooks)
		r.Put("/{id}", bookController.UpdateBook)
		r.Delete("/{id}", bookController.DeleteBook)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Post("/create", authorController.CreateAuthor)
		r.Get("/{id}", authorController.GetAuthor)
		r.Get("/", authorController.GetAllAuthors)
		r.Put("/{id}", authorController.UpdateAuthor)
		r.Delete("/{id}", authorController.DeleteAuthor)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/create", userController.CreateUser)
		r.Get("/{id}", userController.GetUser)
		r.Get("/", userController.GetAllUsers)
		r.Put("/{id}", userController.UpdateUser)
		r.Delete("/{id}", userController.DeleteUser)
	})

	return r
}
