package client_in2_book

import (
	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
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
	// 类别ID
	CategoryKey string `json:"categoryKey"`
	// 代码语言
	CodeLanguage CodeLanguage `json:"codeLanguage"`
	// 简介
	Comment string `json:"comment"`
	// 封面图片key
	CoverKey string `json:"coverKey"`
	// 是否精选
	Selected Bool `json:"selected"`
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

type Bool = github_com_johnnyeven_libtools_courier_enumeration.Bool

type Category struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 业务ID
	CategoryKey string `json:"categoryKey"`
	// 图标类名
	IconClassName string `json:"iconClassName"`
	// 分类名
	Name string `json:"name"`
	// 是否保留为系统预设
	Reserved Bool `json:"reserved"`
	// 排序
	Sort int32 `json:"sort"`
}

type CategoryList []Category

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
	// 书籍分类
	CategoryKey string `json:"categoryKey"`
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

type CreateCategoryBody struct {
	// 分类Key
	CategoryKey string `json:"categoryKey"`
	// 图标类名
	IconClassName string `json:"iconClassName"`
	// 分类名
	Name string `json:"name"`
	// 是否保留为系统预设
	Reserved Bool `default:"FALSE" json:"reserved,omitempty"`
	// 排序
	Sort int32 `default:"0" json:"sort,omitempty"`
}

type CreateTagBody struct {
	// 标签名称
	Name string `json:"name"`
}

type ErrorField = github_com_johnnyeven_libtools_courier_status_error.ErrorField

type ErrorFields = github_com_johnnyeven_libtools_courier_status_error.ErrorFields

type GetBooksByTagResult struct {
	//
	Data BookMetaList `json:"data"`
	//
	Total int32 `json:"total"`
}

type GetBooksMetaResult struct {
	//
	Data BookMetaList `json:"data"`
	//
	Total int32 `json:"total"`
}

type GetCategoriesResult struct {
	//
	Data CategoryList `json:"data"`
	//
	Total int32 `json:"total"`
}

type JSONBytes = github_com_johnnyeven_libtools_courier_swagger.JSONBytes

type MetaItem struct {
	//
	Label string `json:"label"`
	//
	Value string `json:"value"`
}

type MySQLTimestamp = github_com_johnnyeven_libtools_timelib.MySQLTimestamp

type OperateTime = github_com_johnnyeven_libtools_sqlx_presets.OperateTime

type PrimaryID = github_com_johnnyeven_libtools_sqlx_presets.PrimaryID

type SetBookTagBody struct {
	// 分类标识
	TagID uint64 `json:"tagID,string"`
}

type SoftDelete = github_com_johnnyeven_libtools_sqlx_presets.SoftDelete

type StatusError = github_com_johnnyeven_libtools_courier_status_error.StatusError

type Tag struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 热度
	Heat uint32 `json:"heat"`
	// 名称
	Name string `json:"name"`
	// 业务ID
	TagID uint64 `json:"tagID,string"`
}

type TagList []Tag

type UpdateBookMetaParams struct {
	// 文档语言
	BookLanguage BookLanguage `json:"bookLanguage,omitempty"`
	// 代码语言
	CodeLanguage CodeLanguage `json:"codeLanguage,omitempty"`
	// 简介
	Comment string `json:"comment,omitempty"`
	// 封面图片key
	CoverKey string `json:"coverKey,omitempty"`
	// 是否精选
	Selected Bool `json:"selected,omitempty"`
	// 状态
	Status BookStatus `json:"status,omitempty"`
	// 标题
	Title string `json:"title,omitempty"`
}

type UpdateCategoryBody struct {
	// 图标类名
	IconClassName string `json:"iconClassName"`
	// 分类名
	Name string `json:"name"`
	// 是否保留为系统预设
	Reserved Bool `json:"reserved,omitempty"`
	// 排序
	Sort int32 `default:"0" json:"sort,omitempty"`
}
