package usecase

import "time"

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
