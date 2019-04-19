package repos

import (
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(middleware.MiddlewareAuth, ReposGroup{})

type ReposGroup struct {
	courier.EmptyOperator
}

func (ReposGroup) Path() string {
	return "/repos"
}
