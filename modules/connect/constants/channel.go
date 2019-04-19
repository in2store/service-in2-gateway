package constants

import "context"

type RepoChannel interface {
	ChannelID() uint64
	GetRepos(ctx context.Context) ([]Repo, error)
}
