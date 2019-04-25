package channels

import (
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(middleware.MiddlewareAuth, ChannelsGroup{})

type ChannelsGroup struct {
	courier.EmptyOperator
}

func (ChannelsGroup) Path() string {
	return "/channels"
}
