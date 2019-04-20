package client_in2_book

import (
	"fmt"

	github_com_johnnyeven_libtools_courier "github.com/johnnyeven/libtools/courier"
	github_com_johnnyeven_libtools_courier_client "github.com/johnnyeven/libtools/courier/client"
	github_com_johnnyeven_libtools_courier_status_error "github.com/johnnyeven/libtools/courier/status_error"
)

type ClientIn2BookInterface interface {
	CreateBook(req CreateBookRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateBookResponse, err error)
	GetBookMetaByBookID(req GetBookMetaByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookMetaByBookIDResponse, err error)
	GetBookRepoByBookID(req GetBookRepoByBookIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetBookRepoByBookIDResponse, err error)
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
	Body CreateBookResult
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
