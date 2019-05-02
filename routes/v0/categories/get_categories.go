package categories

import (
	"context"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/httplib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetCategories{}))
}

// 获取分类列表
type GetCategories struct {
	httpx.MethodGet
	httplib.Pager
}

func (req GetCategories) Path() string {
	return ""
}

func (req GetCategories) Output(ctx context.Context) (result interface{}, err error) {
	result, err = modules.GetCategories(req.Pager, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetCategories] modules.GetCategories err: %v, request: %+v", err, req)
	}
	return
}
