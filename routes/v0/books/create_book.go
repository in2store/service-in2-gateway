package books

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/courier/status_error"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(middleware.MiddlewareAuth, CreateBook{}))
}

type CreateBookBody struct {
	// 文档语言
	BookLanguage client_in2_book.BookLanguage `json:"bookLanguage" default:""`
	// 代码语言
	CodeLanguage client_in2_book.CodeLanguage `json:"codeLanguage" default:""`
	// 简介
	Comment string `json:"comment" default:""`
	// 封面图片key
	CoverKey string `json:"coverKey" default:""`
	// 标题
	Title string `json:"title"`
	// 书籍分类
	CategoryKey string `json:"categoryKey"`
	// 通道ID
	ChannelID uint64 `json:"channelID,string"`
	// 分支名
	RepoBranchName string `json:"repoBranchName" default:"master"`
	// 代码库全名
	RepoFullName string `json:"repoFullName"`
	// Summary文件相对路径
	SummaryPath string `json:"summaryPath" default:"SUMMARY.md"`
}

// 创建书籍
type CreateBook struct {
	httpx.MethodPost
	Body CreateBookBody `name:"body" in:"body"`
}

func (req CreateBook) Path() string {
	return ""
}

func (req CreateBook) Output(ctx context.Context) (result interface{}, err error) {
	user := middleware.GetAuthUserFromContext(ctx)

	// 获取通道配置
	channel, err := modules.GetChannelByChannelID(req.Body.ChannelID, global.Config.ClientAuth)
	if err != nil {
		logrus.Errorf("[CreateBook] modules.GetChannelByChannelID err: %v, request: %d", err, req.Body.ChannelID)
		return nil, errors.UpstreamError
	}

	request := client_in2_book.CreateBookBody{
		CreateBookMetaParams: client_in2_book.CreateBookMetaParams{
			BookLanguage: req.Body.BookLanguage,
			CodeLanguage: req.Body.CodeLanguage,
			Comment:      req.Body.Comment,
			CoverKey:     req.Body.CoverKey,
			Title:        req.Body.Title,
			CategoryKey:  req.Body.CategoryKey,
			UserID:       user.User.UserID,
		},
		CreateBookRepoParams: client_in2_book.CreateBookRepoParams{
			ChannelID:      req.Body.ChannelID,
			EntryURL:       channel.RawURL,
			RepoBranchName: req.Body.RepoBranchName,
			RepoFullName:   req.Body.RepoFullName,
			SummaryPath:    req.Body.SummaryPath,
		},
	}
	result, err = modules.CreateBook(request, global.Config.ClientBook)
	if err != nil {
		if status_error.FromError(err).Key == "BookConflict" {
			return nil, err
		}
		logrus.Errorf("[CreateBook] err: %v, request: %+v", err, request)
		return nil, errors.UpstreamError
	}
	return
}
