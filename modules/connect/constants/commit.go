package constants

import "time"

type GetCommitsParams struct {
	SHA       string    `json:"sha"`
	Path      string    `json:"path"`
	Author    string    `json:"author"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type Commit interface {
}
