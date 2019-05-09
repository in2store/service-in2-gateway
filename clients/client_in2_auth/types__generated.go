package client_in2_auth

import (
	github_com_johnnyeven_libtools_courier_status_error "github.com/johnnyeven/libtools/courier/status_error"
	github_com_johnnyeven_libtools_courier_swagger "github.com/johnnyeven/libtools/courier/swagger"
	github_com_johnnyeven_libtools_sqlx_presets "github.com/johnnyeven/libtools/sqlx/presets"
	github_com_johnnyeven_libtools_timelib "github.com/johnnyeven/libtools/timelib"
)

type Channel struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 认证URL
	AuthURL string `json:"authURL"`
	// 业务ID
	ChannelID uint64 `json:"channelID,string"`
	// ClientID
	ClientID string `json:"clientID"`
	// ClientSecret
	ClientSecret string `json:"clientSecret"`
	// 名称
	Name string `json:"name"`
	// raw文件访问URL
	RawURL string `json:"rawURL"`
	// 交换tokenURL
	TokenURL string `json:"tokenURL"`
}

type ChannelList []Channel

type CreateChannelParams struct {
	// 认证URL
	AuthURL string `json:"authURL"`
	// ClientID
	ClientID string `json:"clientId"`
	// ClientSecret
	ClientSecret string `json:"clientSecret"`
	// 名称
	Name string `json:"name"`
	// 交换tokenURL
	TokenURL string `json:"tokenURL"`
}

type ErrorField = github_com_johnnyeven_libtools_courier_status_error.ErrorField

type ErrorFields = github_com_johnnyeven_libtools_courier_status_error.ErrorFields

type GetAuthURLResult struct {
	//
	URL string `json:"url"`
}

type GetTokensResult struct {
	//
	Data TokenList `json:"data"`
	//
	Total int32 `json:"total"`
}

type JSONBytes = github_com_johnnyeven_libtools_courier_swagger.JSONBytes

type MySQLTimestamp = github_com_johnnyeven_libtools_timelib.MySQLTimestamp

type OperateTime = github_com_johnnyeven_libtools_sqlx_presets.OperateTime

type PrimaryID = github_com_johnnyeven_libtools_sqlx_presets.PrimaryID

type Session struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// 过期时间
	ExpireTime MySQLTimestamp `json:"expireTime"`
	// 业务ID
	SessionID string `json:"sessionID"`
	// 用户ID
	UserID uint64 `json:"userID,string"`
}

type SoftDelete = github_com_johnnyeven_libtools_sqlx_presets.SoftDelete

type StatusError = github_com_johnnyeven_libtools_courier_status_error.StatusError

type Token struct {
	//
	PrimaryID
	//
	OperateTime
	//
	SoftDelete
	// AccessToken is the token that authorizes and authenticates the requests.
	AccessToken string `json:"accessToken"`
	// 通道ID
	ChannelID uint64 `json:"channelID,string"`
	// Expiry is the optional expiration time of the access token.
	ExpiryTime MySQLTimestamp `json:"expiry"`
	// RefreshToken is a token that's used by the application(as opposed to the user) to refresh the access token if it expires.
	RefreshToken string `json:"refreshToken"`
	// 业务ID
	TokenID uint64 `json:"tokenID,string"`
	// TokenType is the type of token. The Type method returns either this or "Bearer", the default.
	TokenType string `json:"tokenType"`
	// 用户ID
	UserID uint64 `json:"userID,string"`
}

type TokenList []Token
