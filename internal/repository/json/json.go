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

	dataDir := filepath.Join(wd, j.DirName)

	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	if err := os.MkdirAll(j.DirName, os.ModePerm); err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	filePath := filepath.Join(j.DirName, fmt.Sprintf("%s-%d.json", analysis.User.Login, time.Now().Unix()))
	file, err := os.Create(filePath)
	if err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(analysis); err != nil {
		return e.Wrap(whereami.WhereAmI(), err)
	}

	fmt.Printf("Saved user timeline to %s\n", filePath)

	return nil
}
