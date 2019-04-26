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

func (c *GithubChannel) GetRepos(ctx context.Context, user string, size, page int32) ([]constants.Repo, constants.PaginationInfo, error) {
	var listOption *github.RepositoryListOptions
	if size > 0 && page > 0 {
		listOption = &github.RepositoryListOptions{
			ListOptions: github.ListOptions{
				Page:    int(page),
				PerPage: int(size),
			},
		}
	}
	repos, resp, err := c.client.Repositories.List(ctx, user, listOption)
	if err != nil {
		return nil, constants.PaginationInfo{}, err
	}

	result := make([]constants.Repo, 0)
	for _, repo := range repos {
		result = append(result, &GithubRepo{
			*repo,
			c,
		})
	}

	return result, constants.PaginationInfo{
		NextPage:  int32(resp.NextPage),
		PrevPage:  int32(resp.PrevPage),
		FirstPage: int32(resp.FirstPage),
		LastPage:  int32(resp.LastPage),
	}, nil
}

func (c *GithubChannel) GetBranches(ctx context.Context, owner string, repo string, size, page int32) ([]constants.Branch, constants.PaginationInfo, error) {
	var listOption *github.ListOptions
	if size > 0 && page > 0 {
		listOption = &github.ListOptions{
			Page:    int(page),
			PerPage: int(size),
		}
	}
	branches, resp, err := c.client.Repositories.ListBranches(ctx, owner, repo, listOption)
	if err != nil {
		return nil, constants.PaginationInfo{}, err
	}

	result := make([]constants.Branch, 0)
	for _, branch := range branches {
		result = append(result, &GithubBranch{
			*branch,
		})
	}

	return result, constants.PaginationInfo{
		NextPage:  int32(resp.NextPage),
		PrevPage:  int32(resp.PrevPage),
		FirstPage: int32(resp.FirstPage),
		LastPage:  int32(resp.LastPage),
	}, nil
}
