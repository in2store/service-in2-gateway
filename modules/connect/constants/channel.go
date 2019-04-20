package constants

import "context"

type RepoChannel interface {
	ChannelID() uint64
	GetRepos(ctx context.Context, size, page int32) ([]Repo, PaginationInfo, error)
}
