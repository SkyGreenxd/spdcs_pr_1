package usecase

import (
	"context"
	"sort"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/domain"
	"github.com/SkyGreenxd/spdcs_pr_1/internal/repository"
	"github.com/SkyGreenxd/spdcs_pr_1/pkg/e"
)

type GitHubApiUseCase struct {
	gitHubRepo repository.GitHubRepo
	infra      GitHubInfrastructure
	draw       Draw
	Username   string
}

func NewGitHubApiUseCase(gitHubRepo repository.GitHubRepo, infra GitHubInfrastructure, draw Draw, username string) *GitHubApiUseCase {
	return &GitHubApiUseCase{
		gitHubRepo: gitHubRepo,
		infra:      infra,
		draw:       draw,
		Username:   username,
	}
}

func (g *GitHubApiUseCase) AccountCareerAnalysis(ctx context.Context) error {
	const op = "GitHubApiUseCase.AccountCareerAnalysis"

	gitHubAccount, err := g.infra.GetAccount(ctx)
	if err != nil {
		return e.Wrap(op, err)
	}

	gitHubRepositories, err := g.infra.GetRepositories(ctx)
	if err != nil {
		return e.Wrap(op, err)
	}

	yearMap := make(map[int]*domain.YearlyActivity)

	for _, repo := range gitHubRepositories {
		year := repo.CreatedAt.Year()

		if yearMap[year] == nil {
			yearMap[year] = &domain.YearlyActivity{
				Year:          year,
				MainLanguages: make(map[string]int),
			}
		}

		activity := yearMap[year]
		activity.Repositories++
		if repo.Language != "" {
			activity.MainLanguages[repo.Language]++
		}
		activity.AvgStarsPerRepo += float64(repo.StargazersCount)

		commitsCount, err := g.infra.GetCommitsCount(ctx, g.Username, repo.Name)
		if err != nil {
			commitsCount = 0
		}

		activity.Commits += commitsCount
	}

	repositories := make([]domain.YearlyActivity, 0, len(yearMap))

	var (
		years             []int
		commits           []int
		repositoriesCount []int
	)
	for _, act := range yearMap {
		if act.Repositories > 0 {
			act.AvgStarsPerRepo /= float64(act.Repositories)
		}
		repositories = append(repositories, *act)
		repositoriesCount = append(repositoriesCount, act.Repositories)
		years = append(years, act.Year)
		commits = append(commits, act.Commits)
	}

	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].Year > repositories[j].Year
	})

	user := domain.NewUser(gitHubAccount.Login, gitHubAccount.PublicRepos, gitHubAccount.Followers, gitHubAccount.Following, gitHubAccount.CreatedAt)
	analysis := domain.NewCareerAnalysis(*user, repositories)

	if err := g.gitHubRepo.SaveUserTimeline(ctx, *analysis); err != nil {
		return e.Wrap(op, err)
	}

	req := NewDrawReq(user.Login, years, commits, repositoriesCount, AggregateLanguages(analysis.Timeline))

	g.draw.DrawCommitsBar(ctx, *req)

	return nil
}

func AggregateLanguages(timeline []domain.YearlyActivity) map[string]int {
	langTotal := make(map[string]int)

	for _, year := range timeline {
		for lang, count := range year.MainLanguages {
			langTotal[lang] += count
		}
	}

	return langTotal
}
