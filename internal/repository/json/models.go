package json

import "time"

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
