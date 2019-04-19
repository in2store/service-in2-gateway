package channels

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"golang.org/x/oauth2"
)

func NewGithubChannel(ctx context.Context, token client_in2_auth.Token) (constants.RepoChannel, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.AccessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &GithubChannel{
		accessToken: token.AccessToken,
		client:      client,
	}, nil
}

type GithubChannel struct {
	accessToken string
	client      *github.Client
}

func (GithubChannel) ChannelID() uint64 {
	return global.Config.GithubChannelID
}

func (*GithubChannel) GetRepos() ([]constants.Repo, error) {

	panic("implement me")
}
