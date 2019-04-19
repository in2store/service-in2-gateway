package client_in2_auth

import (
	"fmt"

	github_com_johnnyeven_libtools_courier "github.com/johnnyeven/libtools/courier"
	github_com_johnnyeven_libtools_courier_client "github.com/johnnyeven/libtools/courier/client"
	github_com_johnnyeven_libtools_courier_status_error "github.com/johnnyeven/libtools/courier/status_error"
)

type ClientIn2AuthInterface interface {
	Authorize(req AuthorizeRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *AuthorizeResponse, err error)
	CreateChannel(req CreateChannelRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateChannelResponse, err error)
	GetAuthURL(req GetAuthURLRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetAuthURLResponse, err error)
	GetSessionBySessionID(req GetSessionBySessionIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetSessionBySessionIDResponse, err error)
	GetTokens(req GetTokensRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTokensResponse, err error)
}

type ClientIn2Auth struct {
	github_com_johnnyeven_libtools_courier_client.Client
}

func (ClientIn2Auth) MarshalDefaults(v interface{}) {
	if cl, ok := v.(*ClientIn2Auth); ok {
		cl.Name = "in2-auth"
		cl.Client.MarshalDefaults(&cl.Client)
	}
}

func (c ClientIn2Auth) Init() {
	c.CheckService()
}

func (c ClientIn2Auth) CheckService() {
	err := c.Request(c.Name+".Check", "HEAD", "/", nil).
		Do().
		Into(nil)
	statusErr := github_com_johnnyeven_libtools_courier_status_error.FromError(err)
	if statusErr.Code == int64(github_com_johnnyeven_libtools_courier_status_error.RequestTimeout) {
		panic(fmt.Errorf("service %s have some error %s", c.Name, statusErr))
	}
}

type AuthorizeRequest struct {
	//
	Code string `in:"query" name:"code"`
	//
	State string `in:"query" name:"state"`
}

func (c ClientIn2Auth) Authorize(req AuthorizeRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *AuthorizeResponse, err error) {
	resp = &AuthorizeResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".Authorize", "GET", "/in2-auth/v0/authorize", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type AuthorizeResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body []byte
}

type CreateChannelRequest struct {
	//
	Body CreateChannelBody `fmt:"json" in:"body"`
}

func (c ClientIn2Auth) CreateChannel(req CreateChannelRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *CreateChannelResponse, err error) {
	resp = &CreateChannelResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".CreateChannel", "POST", "/in2-auth/v0/channels", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type CreateChannelResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body Channel
}

type GetAuthURLRequest struct {
	// ChannelID
	ChannelID uint64 `in:"path" name:"channelId"`
}

func (c ClientIn2Auth) GetAuthURL(req GetAuthURLRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetAuthURLResponse, err error) {
	resp = &GetAuthURLResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetAuthURL", "GET", "/in2-auth/v0/authorize/:channelId", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetAuthURLResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body GetAuthURLResult
}

type GetSessionBySessionIDRequest struct {
	// SessionID
	SessionID string `in:"path" name:"sessionID"`
}

func (c ClientIn2Auth) GetSessionBySessionID(req GetSessionBySessionIDRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetSessionBySessionIDResponse, err error) {
	resp = &GetSessionBySessionIDResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetSessionBySessionID", "GET", "/in2-auth/v0/sessions/:sessionID", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetSessionBySessionIDResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body Session
}

type GetTokensRequest struct {
	// 用户ID
	UserID uint64 `in:"query" name:"userID"`
	// 通道ID
	ChannelID uint64 `in:"query" name:"channelID"`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `default:"10" in:"query" name:"size,omitempty"`
	// 分页偏移
	// 默认为 0
	Offset int32 `in:"query" name:"offset,omitempty"`
}

func (c ClientIn2Auth) GetTokens(req GetTokensRequest, metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *GetTokensResponse, err error) {
	resp = &GetTokensResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".GetTokens", "GET", "/in2-auth/v0/tokens", req, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type GetTokensResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body GetTokensResult
}

func (c ClientIn2Auth) Swagger(metas ...github_com_johnnyeven_libtools_courier.Metadata) (resp *SwaggerResponse, err error) {
	resp = &SwaggerResponse{}
	resp.Meta = github_com_johnnyeven_libtools_courier.Metadata{}

	err = c.Request(c.Name+".Swagger", "GET", "/in2-auth", nil, metas...).
		Do().
		BindMeta(resp.Meta).
		Into(&resp.Body)

	return
}

type SwaggerResponse struct {
	Meta github_com_johnnyeven_libtools_courier.Metadata
	Body JSONBytes
}
