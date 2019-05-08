package channels

import (
	"github.com/johnnyeven/libtools/courier"
)

var Router = courier.NewRouter(ChannelsGroup{})

type ChannelsGroup struct {
	courier.EmptyOperator
}

func (ChannelsGroup) Path() string {
	return "/channels"
}
