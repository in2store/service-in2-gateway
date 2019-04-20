package repos

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/in2store/service-in2-gateway/modules/connect"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/johnnyeven/libtools/timelib"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(GetRepos{}))
}

// 获取Git库列表
type GetRepos struct {
	httpx.MethodGet
	// 通道ID
	ChannelID uint64 `name:"channelID,string" in:"query"`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `name:"size" in:"query" default:"10"  validate:"@int32[-1,50]"`
	// 页码
	// 默认为 0
	Page int32 `name:"page,omitempty" in:"query" validate:"@int32[0,]"`
}

func (req GetRepos) Path() string {
	return ""
}

type GetReposResultItem struct {
	Name       string                 `json:"name"`
	FullName   string                 `json:"fullName"`
	HtmlUrl    string                 `json:"htmlUrl"`
	CloneUrl   string                 `json:"cloneUrl"`
	CreateTime timelib.MySQLTimestamp `json:"createTime"`
	UpdateTime timelib.MySQLTimestamp `json:"updateTime"`
	PushTime   timelib.MySQLTimestamp `json:"pushTime"`
}

type GetReposResult struct {
	Data []GetReposResultItem `json:"data"`
	constants.PaginationInfo
}

func (req GetRepos) Output(ctx context.Context) (result interface{}, err error) {
	user := middleware.GetAuthUserFromContext(ctx)
	tokenRequest := client_in2_auth.GetTokensRequest{
		UserID:    user.UserID,
		ChannelID: req.ChannelID,
		Size:      -1,
	}
	tokens, err := modules.GetTokens(tokenRequest, global.Config.ClientAuth)
	if err != nil {
		logrus.Errorf("[GetRepos] modules.GetTokens err: %v, request: %+v", err, tokenRequest)
		return nil, errors.UpstreamError
	}
	if tokens.Total == 0 {
		logrus.Errorf("[GetRepos] modules.GetTokens err: not found, request: %+v", tokenRequest)
		return nil, errors.NotFound.StatusError().WithMsg("无法获取到Token").WithErrTalk()
	}

	repoCreator, exist := connect.RepoConnector.GetChannelCreator(req.ChannelID)
	if !exist {
		logrus.Errorf("[GetRepos] connect.RepoConnector.GetChannelCreator err: not found, request: %d", req.ChannelID)
		return nil, errors.InternalError.StatusError().WithMsg("无法获取仓库处理器").WithErrTalk()
	}
	channel, err := repoCreator(ctx, tokens.Data[0])
	if err != nil {
		logrus.Errorf("[GetRepos] RepoCreator err: %v, request: %+v", err, tokens.Data[0])
		return nil, errors.InternalError
	}

	repos, pagination, err := channel.GetRepos(ctx, req.Size, req.Page)
	if err != nil {
		logrus.Errorf("[GetRepos] channel.GetRepos() err: %v", err)
		return nil, errors.UpstreamError.StatusError().WithDesc(err.Error())
	}

	data := make([]GetReposResultItem, 0)
	for _, r := range repos {
		data = append(data, GetReposResultItem{
			Name:       r.GetName(),
			FullName:   r.GetFullName(),
			HtmlUrl:    r.GetHTMLURL(),
			CloneUrl:   r.GetCloneURL(),
			CreateTime: timelib.MySQLTimestamp(r.GetCreatedAt()),
			UpdateTime: timelib.MySQLTimestamp(r.GetUpdatedAt()),
			PushTime:   timelib.MySQLTimestamp(r.GetPushedAt()),
		})
	}

	return GetReposResult{
		Data: data,
		PaginationInfo: constants.PaginationInfo{
			NextPage:  pagination.NextPage,
			PrevPage:  pagination.PrevPage,
			FirstPage: pagination.FirstPage,
			LastPage:  pagination.LastPage,
		},
	}, nil
}
