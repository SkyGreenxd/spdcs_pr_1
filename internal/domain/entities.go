package domain

import "time"

type User struct {
	Login       string
	PublicRepos int
	Followers   int
	Following   int
	CreatedAt   time.Time
}

type YearlyActivity struct {
	Year            int
	Repositories    int
	Commits         int
	MainLanguages   map[string]int
	AvgStarsPerRepo float64
}

type CareerAnalysis struct {
	User     User
	Timeline []YearlyActivity
}

func NewUser(login string, publicRepos int, followers int, following int, createdAt time.Time) *User {
	return &User{
		Login:       login,
		PublicRepos: publicRepos,
		Followers:   followers,
		Following:   following,
		CreatedAt:   createdAt,
	}
}

func NewRepository(year int, repos int, commits int, mainLanguages map[string]int, avgStarsPerRepo float64, externalContribs int) *YearlyActivity {
	return &YearlyActivity{
		Year:            year,
		Repositories:    repos,
		Commits:         commits,
		MainLanguages:   mainLanguages,
		AvgStarsPerRepo: avgStarsPerRepo,
	}
}

func NewCareerAnalysis(user User, timeline []YearlyActivity) *CareerAnalysis {
	return &CareerAnalysis{
		User:     user,
		Timeline: timeline,
	}
}
