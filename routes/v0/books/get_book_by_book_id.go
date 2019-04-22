package books

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func init() {
	Router.Register(courier.NewRouter(GetBookByBookID{}))
}

// 根据书籍ID获取书籍信息
type GetBookByBookID struct {
	httpx.MethodGet
	// 书籍ID
	BookID uint64 `name:"bookID,string" in:"path"`
}

func (req GetBookByBookID) Path() string {
	return "/:bookID"
}

type GetBookByBookIDResult struct {
	// 业务ID
	BookID uint64 `json:"bookID,string"`
	// 文档语言
	BookLanguage client_in2_book.BookLanguage `json:"bookLanguage"`
	// 代码语言
	CodeLanguage client_in2_book.CodeLanguage `json:"codeLanguage"`
	// 简介
	Comment string `json:"comment"`
	// 封面图片key
	CoverKey string `json:"coverKey"`
	// 状态
	Status client_in2_book.BookStatus `json:"status"`
	// 标题
	Title string `json:"title"`
	// 作者ID
	UserID uint64 `json:"userID,string"`
	// 通道ID
	ChannelID uint64 `json:"channelID,string"`
	// 入口地址
	EntryURL string `json:"entryURL"`
	// 代码库分支
	RepoBranchName string `json:"repoBranchName"`
	// 代码库全名
	RepoFullName string `json:"repoFullName"`
	// Summary文件相对地址
	SummaryPath string `json:"summaryPath"`
	// raw文件访问URL
	RawURL string `json:"rawURL"`
	// Summary文件入口
	EntrySummary string `json:"entrySummary"`
}

func (req GetBookByBookID) Output(ctx context.Context) (result interface{}, err error) {
	meta, err := modules.GetBookMetaByBookID(req.BookID, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetBookByBookID] modules.GetBookMetaByBookID err: %v, request: %d", err, req.BookID)
		return nil, errors.UpstreamError
	}
	repo, err := modules.GetBookRepoByBookID(req.BookID, global.Config.ClientBook)
	if err != nil {
		logrus.Errorf("[GetBookByBookID] modules.GetBookRepoByBookID err: %v, request: %d", err, req.BookID)
		return nil, errors.UpstreamError
	}
	channel, err := modules.GetChannelByChannelID(repo.ChannelID, global.Config.ClientAuth)
	if err != nil {
		logrus.Errorf("[GetBookByBookID] modules.GetChannelByChannelID err: %v, request: %d", err, repo.ChannelID)
		return nil, errors.UpstreamError
	}

	return GetBookByBookIDResult{
		BookID:         meta.BookID,
		BookLanguage:   meta.BookLanguage,
		CodeLanguage:   meta.CodeLanguage,
		Comment:        meta.Comment,
		CoverKey:       meta.CoverKey,
		Status:         meta.Status,
		Title:          meta.Title,
		UserID:         meta.UserID,
		ChannelID:      repo.ChannelID,
		EntryURL:       repo.EntryURL,
		RepoBranchName: repo.RepoBranchName,
		RepoFullName:   repo.RepoFullName,
		SummaryPath:    repo.SummaryPath,
		RawURL:         channel.RawURL,
		EntrySummary:   filepath.Join(channel.RawURL, repo.RepoFullName, repo.RepoBranchName, repo.SummaryPath),
	}, nil
}
