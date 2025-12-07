package infrastructure

import (
	"context"
	"os"
	"path/filepath"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/usecase"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type GoEcharts struct{}

func NewGoEcharts() *GoEcharts {
	return &GoEcharts{}
}

func (g *GoEcharts) DrawCommitsBar(ctx context.Context, req usecase.DrawReq) {
	// --- График коммитов ---
	commitBar := charts.NewBar()
	commitBar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Коммиты по годам"}))

	commitItems := make([]opts.BarData, len(req.Commits))
	for i, v := range req.Commits {
		commitItems[i] = opts.BarData{Value: v}
	}
	commitBar.SetXAxis(req.Years).AddSeries("Коммиты", commitItems)

	// --- График репозиториев ---
	repoBar := charts.NewBar()
	repoBar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Репозитории по годам"}))

	repoItems := make([]opts.BarData, len(req.Repos))
	for i, v := range req.Repos {
		repoItems[i] = opts.BarData{Value: v}
	}
	repoBar.SetXAxis(req.Years).AddSeries("Репозитории", repoItems)

	// --- Pie-график языков ---
	langPie := g.buildLanguagesPie(req.Languages)

	// --- Создаем страницу ---
	page := components.NewPage()
	page.AddCharts(commitBar, repoBar, langPie)

	// --- Папка results/<username> ---
	dirPath := filepath.Join("results", req.Username)
	_ = os.MkdirAll(dirPath, os.ModePerm)

	// --- Файл results/<username>/dashboard.html ---
	filePath := filepath.Join(dirPath, "dashboard.html")
	f, _ := os.Create(filePath)
	defer f.Close()

	page.Render(f)
}

func (g *GoEcharts) buildLanguagesPie(langMap map[string]int) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Суммарное использование языков",
	}))

	items := make([]opts.PieData, 0, len(langMap))
	for lang, count := range langMap {
		items = append(items, opts.PieData{
			Name:  lang,
			Value: count,
		})
	}

	pie.AddSeries("Языки", items)
	return pie
}
