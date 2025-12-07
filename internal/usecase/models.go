package usecase

import (
	"time"
)

type UserRes struct {
	Login       string
	PublicRepos int
	Followers   int
	Following   int
	CreatedAt   time.Time
}

type GitHubRes struct {
	Name            string
	Language        string
	StargazersCount int
	ForksCount      int
	CreatedAt       time.Time
}

type DrawReq struct {
	Username  string
	Years     []int
	Commits   []int
	Repos     []int
	Languages map[string]int
}

func NewDrawReq(username string, years []int, commits []int, repositoriesCount []int, languages map[string]int) *DrawReq {
	return &DrawReq{
		Username:  username,
		Years:     years,
		Commits:   commits,
		Repos:     repositoriesCount,
		Languages: languages,
	}
}
