package books

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/clients/client_in2_user"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/enumeration"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/sqlx/presets"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBooksMetaWithDetail{}))
}

// 获取书籍元数据列表（包含用户及标签）
type GetBooksMetaWithDetail struct {
	httpx.MethodGet
	// 用户ID
	UserID uint64 `name:"userID,string" in:"query" default:""`
	// 书籍状态
	Status client_in2_book.BookStatus `name:"status" in:"query" default:""`
	// 分类
	CategoryKey string `in:"query" name:"categoryKey" default:""`
	// 是否精选
	Selected enumeration.Bool `in:"query" name:"selected" default:""`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `name:"size" in:"query" default:"10"  validate:"@int32[-1,10]"`
	// 分页偏移
	// 默认为 0
	Offset int32 `name:"offset,omitempty" in:"query" validate:"@int32[0,]"`
}

type GetBooksMetaWithDetailResult struct {
	Data  []GetBooksMetaWithDetailItem `json:"data"`
	Total int32                        `json:"total"`
}

type GetBooksMetaWithDetailItem struct {
	// 文档ID
	BookID uint64 `json:"bookID,string"`
	// 文档语言
	BookLanguage client_in2_book.BookLanguage `json:"bookLanguage"`
	// 类别ID
	CategoryKey string `json:"categoryKey"`
	// 代码语言
	CodeLanguage client_in2_book.CodeLanguage `json:"codeLanguage"`
	// 简介
	Comment string `json:"comment"`
	// 封面图片key
	CoverKey string `json:"coverKey"`
	// 是否精选
	Selected enumeration.Bool `json:"selected"`
	// 状态
	Status client_in2_book.BookStatus `json:"status"`
	// 标题
	Title string `json:"title"`
	// 标签
	Tags client_in2_book.TagList `json:"tags"`
	// 用户信息
	User *client_in2_user.User `json:"user"`

	presets.OperateTime
}

func (req GetBooksMetaWithDetail) Path() string {
	return "/:bookID/meta-full"
}

func (req GetBooksMetaWithDetail) Output(ctx context.Context) (result interface{}, err error) {
	request := client_in2_book.GetBooksMetaRequest{
		UserID:      req.UserID,
		Status:      req.Status,
		CategoryKey: req.CategoryKey,
		Selected:    req.Selected,
		Size:        req.Size,
		Offset:      req.Offset,
	}
	resp, err := modules.GetBooksMeta(request, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetBooksMeta] modules.GetBooksMeta err: %v, request: %+v", err, request)
		return
	}

	data := make([]GetBooksMetaWithDetailItem, 0)
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
		data = append(data, GetBooksMetaWithDetailItem{
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
	return GetBooksMetaWithDetailResult{
		Data:  data,
		Total: resp.Total,
	}, nil
}
