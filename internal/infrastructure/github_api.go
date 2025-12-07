package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	u "net/url"
	"path"
	"regexp"
	"strconv"

	"github.com/SkyGreenxd/spdcs_pr_1/internal/usecase"
	"github.com/SkyGreenxd/spdcs_pr_1/pkg/e"
	"github.com/jimlawless/whereami"
)

type GitHubClient struct {
	Client   *http.Client
	Username string
}

func NewGitHubClient(httpClient *http.Client, username string) *GitHubClient {
	return &GitHubClient{
		Client:   httpClient,
		Username: username,
	}
}

func (ghc *GitHubClient) GetAccount(ctx context.Context) (*usecase.UserRes, error) {
	const baseURL = "https://api.github.com/users/"
	url, err := u.Parse(baseURL)
	if err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}

	url.Path = path.Join(url.Path, ghc.Username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}

	resp, err := ghc.Client.Do(req)
	if err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, e.Wrap(whereami.WhereAmI(), fmt.Errorf("status %d, body: %s", resp.StatusCode, string(bodyBytes)))
	}

	var infoModel UserRes
	if err := json.NewDecoder(resp.Body).Decode(&infoModel); err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}

	return toUseCaseGetAccountInfo(&infoModel), nil
}

func (ghc *GitHubClient) GetRepositories(ctx context.Context) ([]usecase.GitHubRes, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", ghc.Username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}

	resp, err := ghc.Client.Do(req)
	if err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, e.Wrap(whereami.WhereAmI(), fmt.Errorf("status %d, body: %s", resp.StatusCode, string(bodyBytes)))
	}

	var repos []GitHubRes
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, e.Wrap(whereami.WhereAmI(), err)
	}

	return toArrUseCaseGetRepositories(repos), nil
}

func (g *GitHubClient) GetCommitsCount(ctx context.Context, owner, repo string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?per_page=1", owner, repo)

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	linkHeader := resp.Header.Get("Link")
	if linkHeader == "" {
		if resp.StatusCode == http.StatusOK {
			return 1, nil
		}
		return 0, fmt.Errorf("unexpected response")
	}

	re := regexp.MustCompile(`&page=(\d+)>; rel="last"`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) < 2 {
		return 0, fmt.Errorf("cannot parse Link header")
	}

	count, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, err
	}
	return count, nil
}
