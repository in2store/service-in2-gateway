package v0

import (
	"github.com/in2store/service-in2-gateway/routes/v0/books"
	"github.com/in2store/service-in2-gateway/routes/v0/categories"
	"github.com/in2store/service-in2-gateway/routes/v0/channels"
	"github.com/in2store/service-in2-gateway/routes/v0/metas"
	"github.com/in2store/service-in2-gateway/routes/v0/repos"
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(V0Group{})

func init() {
	Router.Register(repos.Router)
	Router.Register(books.Router)
	Router.Register(channels.Router)
	Router.Register(metas.Router)
	Router.Register(categories.Router)
}

type V0Group struct {
	courier.EmptyOperator
}

func (V0Group) Path() string {
	return "/v0"
}
