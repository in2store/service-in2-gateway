package metas

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetBookMetas{}))
}

// 获取书籍常量配置
type GetBookMetas struct {
	httpx.MethodGet
}

func (req GetBookMetas) Path() string {
	return "/book"
}

type GetBookMetasResult struct {
	BookLanguage []client_in2_book.MetaItem `json:"bookLanguage"`
	CodeLanguage []client_in2_book.MetaItem `json:"codeLanguage"`
}

func (req GetBookMetas) Output(ctx context.Context) (result interface{}, err error) {
	bookLanguage, err := global.Config.ClientBook.GetBookLanguage()
	if err != nil {
		logrus.Warningf("[GetBookMetas] ClientBook.GetBookLanguage err: %v", err)
		return nil, errors.UpstreamError
	}

	codeLanguage, err := global.Config.ClientBook.GetCodeLanguage()
	if err != nil {
		logrus.Warningf("[GetBookMetas] ClientBook.GetCodeLanguage err: %v", err)
		return nil, errors.UpstreamError
	}

	return GetBookMetasResult{
		BookLanguage: bookLanguage.Body,
		CodeLanguage: codeLanguage.Body,
	}, nil
}
