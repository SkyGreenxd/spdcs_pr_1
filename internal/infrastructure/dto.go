package infrastructure

import (
	"time"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/usecase"
)

type UserRes struct {
	Login       string    `json:"login"`
	PublicRepos int       `json:"public_repos"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
}

type GitHubRes struct {
	Name            string    `json:"name"`
	Language        string    `json:"language"`
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	CreatedAt       time.Time `json:"created_at"`
}

func toUseCaseGetAccountInfo(res *UserRes) *usecase.UserRes {
	return &usecase.UserRes{
		Login:       res.Login,
		PublicRepos: res.PublicRepos,
		Followers:   res.Followers,
		Following:   res.Following,
		CreatedAt:   res.CreatedAt,
	}
}

func toUseCaseGetRepositories(res *GitHubRes) *usecase.GitHubRes {
	return &usecase.GitHubRes{
		Name:            res.Name,
		Language:        res.Language,
		StargazersCount: res.StargazersCount,
		ForksCount:      res.ForksCount,
		CreatedAt:       res.CreatedAt,
	}
}

func toArrUseCaseGetRepositories(arr []GitHubRes) []usecase.GitHubRes {
	res := make([]usecase.GitHubRes, len(arr))
	for i, v := range arr {
		res[i] = *toUseCaseGetRepositories(&v)
	}

	return res
}
