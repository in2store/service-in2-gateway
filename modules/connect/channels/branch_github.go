package channels

import (
	"github.com/google/go-github/github"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
)

type GithubBranch struct {
	github.Branch
}

func (b GithubBranch) GetCommit() constants.Commit {
	panic("implement me")
}
