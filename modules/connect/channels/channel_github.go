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

func (c *GithubChannel) GetRepos(ctx context.Context) ([]constants.Repo, error) {
	repos, _, err := c.client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	if err != nil {
		return nil, err
	}

	result := make([]constants.Repo, 0)
	for _, repo := range repos {
		result = append(result, &GithubRepo{
			*repo,
		})
	}

	return result, nil
}
