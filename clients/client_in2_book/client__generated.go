package client_in2_book

import (
	"fmt"

	github_com_johnnyeven_libtools_courier "github.com/johnnyeven/libtools/courier"
	github_com_johnnyeven_libtools_courier_client "github.com/johnnyeven/libtools/courier/client"
	github_com_johnnyeven_libtools_courier_status_error "github.com/johnnyeven/libtools/courier/status_error"
)

type ClientIn2BookInterface interface {
	CreateBook(req CreateBookRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateBookResponse, err error)
	CreateCategory(req CreateCategoryRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateCategoryResponse, err error)
	CreateTag(req CreateTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateTagResponse, err error)
	GetBookLanguage(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookLanguageResponse, err error)
	GetBookMetaByBookID(req GetBookMetaByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookMetaByBookIDResponse, err error)
	GetBookRepoByBookID(req GetBookRepoByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookRepoByBookIDResponse, err error)
	GetBooksByTag(req GetBooksByTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBooksByTagResponse, err error)
	GetBooksMeta(req GetBooksMetaRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBooksMetaResponse, err error)
	GetCategories(req GetCategoriesRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetCategoriesResponse, err error)
	GetCodeLanguage(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetCodeLanguageResponse, err error)
	GetTags(req GetTagsRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTagsResponse, err error)
	GetTagsByBookID(req GetTagsByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTagsByBookIDResponse, err error)
	SetBookTag(req SetBookTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *SetBookTagResponse, err error)
	UpdateBookMeta(req UpdateBookMetaRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *UpdateBookMetaResponse, err error)
	UpdateCategory(req UpdateCategoryRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *UpdateCategoryResponse, err error)
}

type ClientIn2Book struct {
	github_com_johnnyeven_libtools_courier_client.Client
}

func (ClientIn2Book) MarshalDefaults(v interface{}) {
	if cl, ok := v.(*ClientIn2Book); ok {
		cl.Name = "in2-book"
		cl.Client.MarshalDefaults(&cl.Client)
	}
}

func (c ClientIn2Book) Init() {
	c.CheckService()
}

func (c ClientIn2Book) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := github_com_johnnyeven_libtools_courier_status_error.FromError(err)
	if statusErr.Code == int64(github_com_johnnyeven_libtools_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

type CreateBookRequest struct {
	//
	Body CreateBookBody `fmt:"json" in:"body"`
}

func (c ClientIn2Book) CreateBook(req CreateBookRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateBookResponse, err error) {
	resp = &CreateBookResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".CreateBook", "POST", "/in2-book/v0/books", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateBookResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body BookMeta
}

type CreateCategoryRequest struct {
	//
	Body CreateCategoryBody `fmt:"json" in:"body"`
}

func (c ClientIn2Book) CreateCategory(req CreateCategoryRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateCategoryResponse, err error) {
	resp = &CreateCategoryResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".CreateCategory", "POST", "/in2-book/v0/categories", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateCategoryResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body Category
}

type CreateTagRequest struct {
	//
	Body CreateTagBody `fmt:"json" in:"body"`
}

func (c ClientIn2Book) CreateTag(req CreateTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateTagResponse, err error) {
	resp = &CreateTagResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".CreateTag", "POST", "/in2-book/v0/tags", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateTagResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body Tag
}

func (c ClientIn2Book) GetBookLanguage(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookLanguageResponse, err error) {
	resp = &GetBookLanguageResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetBookLanguage", "GET", "/in2-book/v0/metas/book-language", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBookLanguageResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body []MetaItem
}

type GetBookMetaByBookIDRequest struct {
	// 书籍ID
	BookID uint64 `in:"path" name:"bookID"`
}

func (c ClientIn2Book) GetBookMetaByBookID(req GetBookMetaByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookMetaByBookIDResponse, err error) {
	resp = &GetBookMetaByBookIDResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetBookMetaByBookID", "GET", "/in2-book/v0/books/:bookID/meta", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBookMetaByBookIDResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body BookMeta
}

type GetBookRepoByBookIDRequest struct {
	// 书籍ID
	BookID uint64 `in:"path" name:"bookID"`
}

func (c ClientIn2Book) GetBookRepoByBookID(req GetBookRepoByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookRepoByBookIDResponse, err error) {
	resp = &GetBookRepoByBookIDResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetBookRepoByBookID", "GET", "/in2-book/v0/books/:bookID/repo", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBookRepoByBookIDResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body BookRepo
}

type GetBooksByTagRequest struct {
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `default:"10" in:"query" name:"size,omitempty"`
	// 分页偏移
	// 默认为 0
	Offset int32 `in:"query" name:"offset,omitempty"`
	// 标签ID（优先级高）
	TagID uint64 `in:"path" name:"tagID"`
	// 标签名称
	Name string `in:"query" name:"name,omitempty"`
}

func (c ClientIn2Book) GetBooksByTag(req GetBooksByTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBooksByTagResponse, err error) {
	resp = &GetBooksByTagResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetBooksByTag", "GET", "/in2-book/v0/tags/:tagID/books", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBooksByTagResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body GetBooksByTagResult
}

type GetBooksMetaRequest struct {
	// 是否精选
	Selected Bool `in:"query" name:"selected,omitempty"`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `default:"10" in:"query" name:"size,omitempty"`
	// 分页偏移
	// 默认为 0
	Offset int32 `in:"query" name:"offset,omitempty"`
	// 用户ID
	UserID uint64 `in:"query" name:"userID,omitempty"`
	// 状态
	Status BookStatus `in:"query" name:"status,omitempty"`
	// 分类
	CategoryKey string `in:"query" name:"categoryKey,omitempty"`
}

func (c ClientIn2Book) GetBooksMeta(req GetBooksMetaRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBooksMetaResponse, err error) {
	resp = &GetBooksMetaResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetBooksMeta", "GET", "/in2-book/v0/books", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetBooksMetaResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body GetBooksMetaResult
}

type GetCategoriesRequest struct {
	// 是否仅获取非保留分类
	FilterReserved Bool `default:"TRUE" in:"query" name:"filterReserved,omitempty"`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `default:"10" in:"query" name:"size,omitempty"`
	// 分页偏移
	// 默认为 0
	Offset int32 `in:"query" name:"offset,omitempty"`
}

func (c ClientIn2Book) GetCategories(req GetCategoriesRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetCategoriesResponse, err error) {
	resp = &GetCategoriesResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetCategories", "GET", "/in2-book/v0/categories", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetCategoriesResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body GetCategoriesResult
}

func (c ClientIn2Book) GetCodeLanguage(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetCodeLanguageResponse, err error) {
	resp = &GetCodeLanguageResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetCodeLanguage", "GET", "/in2-book/v0/metas/code-language", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetCodeLanguageResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body []MetaItem
}

type GetTagsRequest struct {
	// 是否依照热度排序
	OrderByHeat Bool `in:"query" name:"orderByHeat"`
	// 是否过滤零热度项
	FilterZeroHeat Bool `in:"query" name:"filterZeroHeat"`
}

func (c ClientIn2Book) GetTags(req GetTagsRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTagsResponse, err error) {
	resp = &GetTagsResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetTags", "GET", "/in2-book/v0/tags", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetTagsResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body TagList
}

type GetTagsByBookIDRequest struct {
	// 文档ID
	BookID uint64 `in:"path" name:"bookID"`
}

func (c ClientIn2Book) GetTagsByBookID(req GetTagsByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTagsByBookIDResponse, err error) {
	resp = &GetTagsByBookIDResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetTagsByBookID", "GET", "/in2-book/v0/books/:bookID/tags", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetTagsByBookIDResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body TagList
}

type SetBookTagRequest struct {
	// 文档ID
	BookID uint64 `in:"path" name:"bookID"`
	//
	Body SetBookTagBody `fmt:"json" in:"body"`
}

func (c ClientIn2Book) SetBookTag(req SetBookTagRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *SetBookTagResponse, err error) {
	resp = &SetBookTagResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".SetBookTag", "POST", "/in2-book/v0/books/:bookID/tags", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SetBookTagResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body []byte
}

func (c ClientIn2Book) Swagger(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *SwaggerResponse, err error) {
	resp = &SwaggerResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".Swagger", "GET", "/in2-book", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body JSONBytes
}

type UpdateBookMetaRequest struct {
	// 书籍ID
	BookID uint64 `in:"path" name:"bookID"`
	//
	Body UpdateBookMetaParams `fmt:"json" in:"body"`
}

func (c ClientIn2Book) UpdateBookMeta(req UpdateBookMetaRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *UpdateBookMetaResponse, err error) {
	resp = &UpdateBookMetaResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".UpdateBookMeta", "PATCH", "/in2-book/v0/books/:bookID/meta", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UpdateBookMetaResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body BookMeta
}

type UpdateCategoryRequest struct {
	// 分类标识
	CategoryKey string `in:"path" name:"categoryKey"`
	//
	Body UpdateCategoryBody `fmt:"json" in:"body"`
}

func (c ClientIn2Book) UpdateCategory(req UpdateCategoryRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *UpdateCategoryResponse, err error) {
	resp = &UpdateCategoryResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".UpdateCategory", "PATCH", "/in2-book/v0/categories/:categoryKey", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type UpdateCategoryResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body []byte
}
