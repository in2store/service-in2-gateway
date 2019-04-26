package constants

import (
	"context"
	"time"
)

type Repo interface {
	GetBranches(ctx context.Context, size, page int32) ([]Branch, PaginationInfo, error)
	GetCommits(req GetCommitsParams) ([]Commit, error)

	GetAllowMergeCommit() bool
	GetAllowRebaseMerge() bool
	GetAllowSquashMerge() bool
	GetArchived() bool
	GetArchiveURL() string
	GetAssigneesURL() string
	GetAutoInit() bool
	GetBlobsURL() string
	GetBranchesURL() string
	GetCloneURL() string
	GetCollaboratorsURL() string
	GetCommentsURL() string
	GetCommitsURL() string
	GetCompareURL() string
	GetContentsURL() string
	GetContributorsURL() string
	GetCreatedAt() time.Time
	GetDefaultBranch() string
	GetDeploymentsURL() string
	GetDescription() string
	GetDownloadsURL() string
	GetEventsURL() string
	GetFork() bool
	GetForksCount() int
	GetForksURL() string
	GetFullName() string
	GetGitCommitsURL() string
	GetGitignoreTemplate() string
	GetGitRefsURL() string
	GetGitTagsURL() string
	GetGitURL() string
	GetHasDownloads() bool
	GetHasIssues() bool
	GetHasPages() bool
	GetHasProjects() bool
	GetHasWiki() bool
	GetHomepage() string
	GetHooksURL() string
	GetHTMLURL() string
	GetID() int64
	GetIssueCommentURL() string
	GetIssueEventsURL() string
	GetIssuesURL() string
	GetKeysURL() string
	GetLabelsURL() string
	GetLanguage() string
	GetLanguagesURL() string
	GetMasterBranch() string
	GetMergesURL() string
	GetMilestonesURL() string
	GetMirrorURL() string
	GetName() string
	GetNetworkCount() int
	GetNodeID() string
	GetNotificationsURL() string
	GetOpenIssuesCount() int
	GetPermissions() map[string]bool
	GetPrivate() bool
	GetPullsURL() string
	GetPushedAt() time.Time
	GetReleasesURL() string
	GetSize() int
	GetSSHURL() string
	GetStargazersCount() int
	GetStargazersURL() string
	GetStatusesURL() string
	GetSubscribersCount() int
	GetSubscribersURL() string
	GetSubscriptionURL() string
	GetSVNURL() string
	GetTagsURL() string
	GetTeamID() int64
	GetTeamsURL() string
	GetTreesURL() string
	GetUpdatedAt() time.Time
	GetURL() string
	GetWatchersCount() int
}
