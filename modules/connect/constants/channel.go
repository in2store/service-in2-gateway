package constants

type RepoChannel interface {
	ChannelID() uint64
	GetRepos() ([]Repo, error)
}
