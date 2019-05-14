package modules

import (
	"github.com/in2store/service-in2-gateway/clients/client_in2_book"
	"github.com/johnnyeven/libtools/httplib"
)

func CreateBook(req client_in2_book.CreateBookBody, client *client_in2_book.ClientIn2Book) (*client_in2_book.BookMeta, error) {
	request := client_in2_book.CreateBookRequest{
		Body: req,
	}
	resp, err := client.CreateBook(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetBookMetaByBookID(bookID uint64, client *client_in2_book.ClientIn2Book) (*client_in2_book.BookMeta, error) {
	request := client_in2_book.GetBookMetaByBookIDRequest{
		BookID: bookID,
	}
	resp, err := client.GetBookMetaByBookID(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetBookRepoByBookID(bookID uint64, client *client_in2_book.ClientIn2Book) (*client_in2_book.BookRepo, error) {
	request := client_in2_book.GetBookRepoByBookIDRequest{
		BookID: bookID,
	}
	resp, err := client.GetBookRepoByBookID(request)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetBooksMeta(req client_in2_book.GetBooksMetaRequest, client *client_in2_book.ClientIn2Book) (*client_in2_book.GetBooksMetaResult, error) {
	resp, err := client.GetBooksMeta(req)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

type GetBooksResult struct {
	Data  []GetBooksResultItem `json:"data"`
	Total int32                `json:"total"`
}

type GetBooksResultItem struct {
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

	client_in2_book.OperateTime
}

func GetBooks(req client_in2_book.GetBooksMetaRequest, client *client_in2_book.ClientIn2Book) (*GetBooksResult, error) {
	resp, err := client.GetBooksMeta(req)
	if err != nil {
		return nil, err
	}

	result := make([]GetBooksResultItem, 0)
	for _, meta := range resp.Body.Data {
		repo, err := GetBookRepoByBookID(meta.BookID, client)
		if err != nil {
			return nil, err
		}
		result = append(result, GetBooksResultItem{
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
			OperateTime: client_in2_book.OperateTime{
				CreateTime: meta.CreateTime,
				UpdateTime: meta.UpdateTime,
			},
		})
	}

	return &GetBooksResult{
		result,
		resp.Body.Total,
	}, nil
}

func GetCategories(pager httplib.Pager, client *client_in2_book.ClientIn2Book) (*client_in2_book.GetCategoriesResult, error) {
	resp, err := client.GetCategories(client_in2_book.GetCategoriesRequest{
		Size:   pager.Size,
		Offset: pager.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

func GetTagsByBookID(bookID uint64, client *client_in2_book.ClientIn2Book) (client_in2_book.TagList, error) {
	resp, err := client.GetTagsByBookID(client_in2_book.GetTagsByBookIDRequest{
		BookID: bookID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func CreateTag(req client_in2_book.CreateTagBody, client *client_in2_book.ClientIn2Book) (*client_in2_book.Tag, error) {
	resp, err := client.CreateTag(client_in2_book.CreateTagRequest{
		Body: req,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Body, err
}

func GetTags(req client_in2_book.GetTagsRequest, client *client_in2_book.ClientIn2Book) (client_in2_book.TagList, error) {
	resp, err := client.GetTags(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func GetBooksByTag(req client_in2_book.GetBooksByTagRequest, client *client_in2_book.ClientIn2Book) (*client_in2_book.GetBooksByTagResult, error) {
	resp, err := client.GetBooksByTag(req)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}
