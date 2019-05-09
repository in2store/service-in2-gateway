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
	Router.Register(courier.NewRouter(Authorize{}))
}

// 处理认证回调
type Authorize struct {
	httpx.MethodGet
	Code  string `name:"code" in:"query"`
	State string `name:"state" in:"query"`
}

func (req Authorize) Path() string {
	return "/:channelID/authorize"
}

func (req Authorize) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_auth.AuthorizeRequest{
		Code:  req.Code,
		State: req.State,
	}
	resp, err := global.Config.ClientAuth.Authorize(request)
	if err != nil {
		logrus.Errorf("[Authorize] ClientAuth.Authorize err: %v, request: %+v", err, req)
		return nil, err
	}
	return resp.Body, nil
}
