package books

import (
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(middleware.MiddlewareAuth, BooksGroup{})

type BooksGroup struct {
	courier.EmptyOperator
}

func (BooksGroup) Path() string {
	return "/books"
}
