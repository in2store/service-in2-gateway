package tags

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/courier/status_error"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(CreateTag{}))
}

// 创建标签
type CreateTag struct {
	httpx.MethodPost
	Body client_in2_book.CreateTagBody `name:"body" in:"body"`
}

func (req CreateTag) Path() string {
	return ""
}

func (req CreateTag) Output(ctx context.Context) (result interface{}, err error) {
	resp, err := modules.CreateTag(req.Body, global.Config.ClientBook)
	if err != nil {
		if status_error.FromError(err).Key == "TagConflict" {
			return nil, err
		}
		logrus.Errorf("[CreateTag] modules.CreateTag err: %v, request: %+v", err, req.Body)
		return nil, errors.UpstreamError
	}
	return resp, nil
}
