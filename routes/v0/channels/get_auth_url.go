package channels

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetAuthURL{}))
}

// 获取认证界面URL
type GetAuthURL struct {
	httpx.MethodGet
	// ChannelID
	ChannelID uint64 `name:"channelID,string" in:"path"`
}

func (req GetAuthURL) Path() string {
	return "/:channelID/auth-url"
}

func (req GetAuthURL) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_auth.GetAuthURLRequest{
		ChannelID: req.ChannelID,
	}
	resp, err := global.Config.ClientAuth.GetAuthURL(request)
	if err != nil {
		logrus.Errorf("[GetAuthURL] ClientAuth.GetAuthURL err: %v, request: %+v", err, request)
		return nil, err
	}
	return resp.Body, nil
}
