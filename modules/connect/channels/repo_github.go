package channels

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"time"
)

type GithubRepo struct {
	github.Repository
	channel *GithubChannel
}

func (repo *GithubRepo) GetBranches(ctx context.Context, size, page int32) ([]constants.Branch, constants.PaginationInfo, error) {
	return repo.channel.GetBranches(ctx, repo.GetOwner().String(), repo.GetFullName(), size, page)
}

func (repo *GithubRepo) GetCommits(req constants.GetCommitsParams) ([]constants.Commit, error) {
	panic("implement me")
}

func (repo *GithubRepo) GetCreatedAt() time.Time {
	return repo.Repository.GetCreatedAt().Time
}

func (repo *GithubRepo) GetPushedAt() time.Time {
	return repo.Repository.GetPushedAt().Time
}

func (repo *GithubRepo) GetUpdatedAt() time.Time {
	return repo.Repository.GetUpdatedAt().Time
}
