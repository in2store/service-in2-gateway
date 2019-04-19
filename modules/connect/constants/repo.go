package constants

type Repo interface {
	GetBranches() ([]Branch, error)
	GetCommits(req GetCommitsParams) ([]Commit, error)
}
