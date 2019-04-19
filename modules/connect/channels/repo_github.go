package channels

import (
	"github.com/google/go-github/github"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"time"
)

type GithubRepo struct {
	github.Repository
}

func (repo *GithubRepo) GetBranches() ([]constants.Branch, error) {
	panic("implement me")
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
