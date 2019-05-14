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
	"github.com/johnnyeven/libtools/httplib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksByTag{}))
}

// 通过标签名称获取书籍列表
type GetBooksByTag struct {
	httpx.MethodGet
	// 标签名称
	Name string `name:"name" in:"query"`
	httplib.Pager
}

func (req GetBooksByTag) Path() string {
	return "/:tagID/books"
}

func (req GetBooksByTag) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_book.GetBooksByTagRequest{
		Size:   req.Size,
		Offset: req.Offset,
		TagID:  0,
		Name:   req.Name,
	}
	result, err = modules.GetBooksByTag(request, global.Config.ClientBook)
	if err != nil {
		if status_error.FromError(err).Key == "TagNotFound" {
			return nil, err
		}
		logrus.Errorf("[GetBooksByTag] modules.GetBooksByTag err: %v, request: %+v", err, req)
		return nil, errors.UpstreamError
	}
	return
}
