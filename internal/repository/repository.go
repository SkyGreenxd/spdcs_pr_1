package repository

import (
	"context"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/domain"
)

type GitHubRepo interface {
	SaveUserTimeline(ctx context.Context, analysis domain.CareerAnalysis) error
}
