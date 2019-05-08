package channels

import (
	"context"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetChannels{}))
}

// 获取通道列表
type GetChannels struct {
	httpx.MethodGet
}

func (req GetChannels) Path() string {
	return ""
}

func (req GetChannels) Output(ctx context.Context) (result interface{}, err error) {
	result, err = modules.GetChannels(global.Config.ClientAuth)
	if err != nil {
		logrus.Errorf("[GetChannels] modules.GetChannels err: %v", err)
		return nil, errors.UpstreamError
	}
	return
}
