package app

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/config"
	"github.com/SkyGreenxd/spdcs_pr_1/internal/infrastructure"
	"github.com/SkyGreenxd/spdcs_pr_1/internal/repository/json"
	"github.com/SkyGreenxd/spdcs_pr_1/internal/usecase"
	"github.com/SkyGreenxd/spdcs_pr_1/pkg/e"
	"github.com/jimlawless/whereami"
)

func Run() {
	const dirName = "results"

	username, err := ReadUsername()
	if err != nil {
		log.Fatal(err)
	}

	creator := json.NewJSONCreator(dirName)

	if err := config.LoadEnv(); err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("TOKEN")

	infra := infrastructure.NewGitHubClient(&http.Client{}, username, token)
	draw := infrastructure.NewGoEcharts()
	ghUsecase := usecase.NewGitHubApiUseCase(creator, infra, draw, username)

	if err := ghUsecase.AccountCareerAnalysis(context.Background()); err != nil {
		log.Fatal(err)
	}

	return
}

func ReadUsername() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя пользователя GitHub: ")

	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return "", e.Wrap(whereami.WhereAmI(), err)
	}

	username = strings.TrimSpace(username)
	fmt.Println("Вы ввели пользователя:", username)

	return username, nil
}
