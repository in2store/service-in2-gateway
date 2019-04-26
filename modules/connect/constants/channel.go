package constants

import "context"

type RepoChannel interface {
	ChannelID() uint64
	GetRepos(ctx context.Context, user string, size, page int32) ([]Repo, PaginationInfo, error)
	GetBranches(ctx context.Context, owner string, repo string, size, page int32) ([]Branch, PaginationInfo, error)
}
