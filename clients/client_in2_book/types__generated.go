package client_in2_book

import (
	github_com_johnnyeven_libtools_courier_status_error "github.com/johnnyeven/libtools/courier/status_error"
	github_com_johnnyeven_libtools_courier_swagger "github.com/johnnyeven/libtools/courier/swagger"
	github_com_johnnyeven_libtools_sqlx_presets "github.com/johnnyeven/libtools/sqlx/presets"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

type BookLanguage = In2BookBookLanguage

type BookMeta struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 业务ID
	BookID uint64 `json:"bookID,string"`
	// 文档语言
	BookLanguage BookLanguage `json:"bookLanguage"`
	// 代码语言
	CodeLanguage CodeLanguage `json:"codeLanguage"`
	// 简介
	Comment string `json:"comment"`
	// 封面图片key
	CoverKey string `json:"coverKey"`
	// 状态
	Status BookStatus `json:"status"`
	// 标题
	Title string `json:"title"`
	// 作者ID
	UserID uint64 `json:"userID,string"`
}

type BookMetaList []BookMeta

type BookRepo struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 书籍ID
	BookID uint64 `json:"bookID,string"`
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
}

type BookStatus = In2BookBookStatus

type CodeLanguage = In2BookCodeLanguage

type CreateBookBody struct {
	//
	CreateBookMetaParams
	//
	CreateBookRepoParams
}

type CreateBookMetaParams struct {
	// 文档语言
	BookLanguage BookLanguage `json:"bookLanguage,omitempty"`
	// 代码语言
	CodeLanguage CodeLanguage `json:"codeLanguage,omitempty"`
	// 简介
	Comment string `json:"comment,omitempty"`
	// 封面图片key
	CoverKey string `json:"coverKey,omitempty"`
	// 标题
	Title string `json:"title"`
	// 作者ID
	UserID uint64 `json:"userID,string"`
}

type CreateBookRepoParams struct {
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
}

type CreateBookResult struct {
	//
	BookID uint64 `json:"bookID,string"`
}

type ErrorField = github_com_johnnyeven_libtools_courier_status_error.ErrorField

type ErrorFields = github_com_johnnyeven_libtools_courier_status_error.ErrorFields

type GetBooksMetaResult struct {
	//
	Data BookMetaList `json:"data"`
	//
	Total int32 `json:"total"`
}

type JSONBytes = github_com_johnnyeven_libtools_courier_swagger.JSONBytes

type MySQLTimestamp = github_com_johnnyeven_libtools_timelib.MySQLTimestamp

type OperateTime = github_com_johnnyeven_libtools_sqlx_presets.OperateTime

type PrimaryID = github_com_johnnyeven_libtools_sqlx_presets.PrimaryID

type SoftDelete = github_com_johnnyeven_libtools_sqlx_presets.SoftDelete

type StatusError = github_com_johnnyeven_libtools_courier_status_error.StatusError
