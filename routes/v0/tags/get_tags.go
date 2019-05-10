package tags

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetTags{}))
}

// 获取标签列表
type GetTags struct {
	httpx.MethodGet
}

func (req GetTags) Path() string {
	return ""
}

func (req GetTags) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_book.GetTagsRequest{
		FilterZeroHeat: enumeration.BOOL__TRUE,
		OrderByHeat:    enumeration.BOOL__TRUE,
	}
	resp, err := modules.GetTags(request, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetTags] modules.GetTags err: %v, request: %+v", err, request)
		return nil, errors.UpstreamError
	}
	return resp, nil
}
