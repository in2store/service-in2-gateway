package repos

import (
	"context"
	"fmt"
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules"
	"github.com/in2store/service-in2-gateway/modules/connect"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
	"github.com/sirupsen/logrus"
	"strings"
)

func init() {
	Router.Register(courier.NewRouter(GetBranches{}))
}

// 获取Git库列表
type GetBranches struct {
	httpx.MethodGet
	// 通道ID
	ChannelID uint64 `name:"channelID,string" in:"query"`
	// 仓库全名
	RepoFullName string `name:"repoFullName" in:"query"`
	// 分页大小
	// 默认为 10，-1 为查询所有
	Size int32 `name:"size" in:"query" default:"10"  validate:"@int32[-1,50]"`
	// 页码
	// 默认为 0
	Page int32 `name:"page" in:"query" default:"1" validate:"@int32[0,]"`
}

func (req GetBranches) Path() string {
	return "/0/branches"
}

type GetBranchesResultItem struct {
	Name      string `json:"name"`
	Protected bool   `json:"protected"`
}

type GetBranchesResult struct {
	Data []GetBranchesResultItem `json:"data"`
	constants.PaginationInfo
}

func (req GetBranches) Validate() ([]string, error) {
	repoPath := strings.Split(req.RepoFullName, "/")
	if len(repoPath) < 2 {
		err := fmt.Errorf("not enough '/'")
		return nil, err
	}
	return repoPath, nil
}

func (req GetBranches) Output(ctx context.Context) (result interface{}, err error) {
	repoPath, err := req.Validate()
	if err != nil {
		logrus.Errorf("[GetBranches] ")
	}

	user := middleware.GetAuthUserFromContext(ctx)
	tokenRequest := client_in2_auth.GetTokensRequest{
		UserID:    user.User.UserID,
		ChannelID: req.ChannelID,
		Size:      -1,
	}
	tokens, err := modules.GetTokens(tokenRequest, global.Config.ClientAuth)
	if err != nil {
		logrus.Errorf("[GetBranches] modules.GetTokens err: %v, request: %+v", err, tokenRequest)
		return nil, errors.UpstreamError
	}
	if tokens.Total == 0 {
		logrus.Errorf("[GetBranches] modules.GetTokens err: not found, request: %+v", tokenRequest)
		return nil, errors.NotFound.StatusError().WithMsg("无法获取到Token").WithErrTalk()
	}

	repoCreator, exist := connect.RepoConnector.GetChannelCreator(req.ChannelID)
	if !exist {
		logrus.Errorf("[GetBranches] connect.RepoConnector.GetChannelCreator err: not found, request: %d", req.ChannelID)
		return nil, errors.InternalError.StatusError().WithMsg("无法获取仓库处理器").WithErrTalk()
	}
	channel, err := repoCreator(ctx, tokens.Data[0])
	if err != nil {
		logrus.Errorf("[GetBranches] RepoCreator err: %v, request: %+v", err, tokens.Data[0])
		return nil, errors.InternalError
	}

	branches, pagination, err := channel.GetBranches(ctx, repoPath[0], repoPath[1], req.Size, req.Page)
	if err != nil {
		logrus.Errorf("[GetBranches] channel.GetBranches() err: %v", err)
		return nil, errors.UpstreamError.StatusError().WithDesc(err.Error())
	}

	data := make([]GetBranchesResultItem, 0)
	for _, b := range branches {
		data = append(data, GetBranchesResultItem{
			Name:      b.GetName(),
			Protected: b.GetProtected(),
		})
	}

	return GetBranchesResult{
		Data: data,
		PaginationInfo: constants.PaginationInfo{
			NextPage:  pagination.NextPage,
			PrevPage:  pagination.PrevPage,
			FirstPage: pagination.FirstPage,
			LastPage:  pagination.LastPage,
		},
	}, nil
}
