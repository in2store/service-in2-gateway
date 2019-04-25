package books

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/httplib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksMeta{}))
}

// 获取书籍元数据列表
type GetBooksMeta struct {
	httpx.MethodGet
	// 书籍状态
	Status client_in2_book.BookStatus `name:"status" in:"query" default:""`
	httplib.Pager
}

func (req GetBooksMeta) Path() string {
	return ""
}

func (req GetBooksMeta) Output(ctx context.Context) (result interface{}, err error) {
	user := middleware.GetAuthUserFromContext(ctx)
	request := client_in2_book.GetBooksMetaRequest{
		UserID: user.UserID,
		Status: req.Status,
		Size:   req.Size,
		Offset: req.Offset,
	}
	result, err = modules.GetBooksMeta(request, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetBooksMeta] modules.GetBooksMeta err: %v, request: %+v", err, request)
	}
	return
}
