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
	"github.com/johnnyeven/libtools/sqlx/presets"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksByTagWithDetail{}))
}

// 通过标签名称获取书籍列表（包含标签及用户信息）
type GetBooksByTagWithDetail struct {
	httpx.MethodGet
	// 标签名称
	Name string `name:"name" in:"query"`
	httplib.Pager
}

func (req GetBooksByTagWithDetail) Path() string {
	return "/:tagID/books-full"
}

func (req GetBooksByTagWithDetail) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_book.GetBooksByTagRequest{
		Size:   req.Size,
		Offset: req.Offset,
		TagID:  0,
		Name:   req.Name,
	}
	resp, err := modules.GetBooksByTag(request, global.Config.ClientBook)
	if err != nil {
		if status_error.FromError(err).Key == "TagNotFound" {
			return nil, err
		}
		logrus.Errorf("[GetBooksByTag] modules.GetBooksByTag err: %v, request: %+v", err, req)
		return nil, errors.UpstreamError
	}

	data := make([]modules.GetBooksMetaWithDetailItem, 0)
	for _, meta := range resp.Data {
		tags, err := modules.GetTagsByBookID(meta.BookID, global.Config.ClientBook)
		if err != nil {
			logrus.Errorf("[GetBooksMeta] modules.GetTagsByBookID err: %v, request: %+d", err, meta.BookID)
			return nil, err
		}
		user, err := modules.GetUserByUserID(meta.UserID, global.Config.ClientUser)
		if err != nil {
			logrus.Errorf("[GetBooksMeta] modules.GetUserByUserID err: %v, request: %+d", err, meta.UserID)
			return nil, err
		}
		data = append(data, modules.GetBooksMetaWithDetailItem{
			BookID:       meta.BookID,
			BookLanguage: meta.BookLanguage,
			CategoryKey:  meta.CategoryKey,
			CodeLanguage: meta.CodeLanguage,
			Comment:      meta.Comment,
			CoverKey:     meta.CoverKey,
			Selected:     meta.Selected,
			Status:       meta.Status,
			Title:        meta.Title,
			Tags:         tags,
			User:         user,
			OperateTime: presets.OperateTime{
				CreateTime: meta.CreateTime,
				UpdateTime: meta.UpdateTime,
			},
		})
	}
	return modules.GetBooksMetaWithDetailResult{
		Data:  data,
		Total: resp.Total,
	}, nil
	return
}
