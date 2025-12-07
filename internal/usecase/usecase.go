package usecase

import (
	"context"
)

type GitHubApiUC interface {
	AccountCareerAnalysis(ctx context.Context) error
}

type GitHubInfrastructure interface {
	GetAccount(ctx context.Context) (*UserRes, error)
	GetRepositories(ctx context.Context) ([]GitHubRes, error)
	GetCommitsCount(ctx context.Context, owner, repo string) (int, error)
}

type Draw interface {
	DrawCommitsBar(ctx context.Context, req DrawReq)
}
