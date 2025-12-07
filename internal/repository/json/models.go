package json

import (
	"time"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/domain"
)

type User struct {
	Login       string    `json:"login"`
	PublicRepos int       `json:"public_repos"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
}

type YearlyActivity struct {
	Year            int            `json:"year"`
	Repositories    int            `json:"repositories"`
	Commits         int            `json:"commits"`
	MainLanguages   map[string]int `json:"main_languages"`
	AvgStarsPerRepo float64        `json:"avg_stars_per_repo"`
}

type CareerAnalysis struct {
	User     User             `json:"user"`
	Timeline []YearlyActivity `json:"timeline"`
}

func toRepoUser(user domain.User) User {
	return User{
		Login:       user.Login,
		PublicRepos: user.PublicRepos,
		Followers:   user.Followers,
		Following:   user.Following,
		CreatedAt:   user.CreatedAt,
	}
}

func toRepoYearlyActivity(y domain.YearlyActivity) YearlyActivity {
	return YearlyActivity{
		Year:            y.Year,
		Repositories:    y.Repositories,
		Commits:         y.Commits,
		MainLanguages:   y.MainLanguages,
		AvgStarsPerRepo: y.AvgStarsPerRepo,
	}
}

func toArrRepoYearlyActivity(arr []domain.YearlyActivity) []YearlyActivity {
	res := make([]YearlyActivity, len(arr))
	for i, y := range arr {
		res[i] = toRepoYearlyActivity(y)
	}

	return res
}

func toCareerAnalysis(user User, y []YearlyActivity) CareerAnalysis {
	return CareerAnalysis{
		User:     user,
		Timeline: y,
	}
}
