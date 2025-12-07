package json

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/domain"
	"github.com/SkyGreenxd/spdcs_pr_1/pkg/e"
	"github.com/jimlawless/whereami"
)

type JSONCreator struct {
	DirName string
}

func NewJSONCreator(dirName string) *JSONCreator {
	return &JSONCreator{
		DirName: dirName,
	}
}

func (j *JSONCreator) SaveUserTimeline(ctx context.Context, analysis domain.CareerAnalysis) error {
	wd, err := os.Getwd()
	if err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	userDir := filepath.Join(wd, j.DirName, analysis.User.Login)

	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	filePath := filepath.Join(
		userDir,
		fmt.Sprintf("%s-%d.json", analysis.User.Login, time.Now().Unix()),
	)

	file, err := os.Create(filePath)
	if err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(toCareerAnalysis(toRepoUser(analysis.User), toArrRepoYearlyActivity(analysis.Timeline))); err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	fmt.Printf("Saved user timeline to %s\n", filePath)

	return nil
}
