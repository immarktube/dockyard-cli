package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

type prRequest struct {
	Title string `json:"title"`
	Head  string `json:"head"`
	Base  string `json:"base"`
	Body  string `json:"body"`
}

func CreatePullRequest(repo config.Repository, title, body string) error {
	branch := getBranch(repo)

	exists, err := PRExists(repo, branch)
	if err != nil {
		return err
	}
	if exists {
		fmt.Printf("🔁 PR already exists for %s [%s]\n", repo.Name, branch)
		return nil
	}

	pr := prRequest{
		Title: title,
		Head:  branch,
		Base:  "master",
		Body:  body,
	}

	prURL, err := doCreatePRRequest(repo, pr)
	if err != nil {
		return err
	}

	fmt.Printf("✅ PR created successfully for %s [%s]: %s\n", repo.Name, branch, prURL)
	return nil
}

func getBranch(repo config.Repository) string {
	if repo.Branch != "" {
		return repo.Branch
	}
	b, err := GetCurrentBranch(repo.Path)
	if err != nil {
		return "master"
	}
	return b
}

func getPullsURL(repo config.Repository) string {
	apiBase := repo.APIBaseURL
	if apiBase == "" {
		apiBase = "https://api.github.com"
	}
	return fmt.Sprintf("%s/repos/%s/%s/pulls", strings.TrimRight(apiBase, "/"), repo.Owner, repo.Name)
}

func doCreatePRRequest(repo config.Repository, pr prRequest) (string, error) {
	url := getPullsURL(repo)
	data, err := json.Marshal(pr)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "token "+repo.AuthToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Creating PR: repo=%s/%s head=%s base=%s\n", repo.Owner, repo.Name, pr.Head, pr.Base)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("GitHub API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var prResp struct {
		HTMLURL string `json:"html_url"`
	}
	if err := json.Unmarshal(bodyBytes, &prResp); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return prResp.HTMLURL, nil
}

func GetCurrentBranch(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func PRExists(repo config.Repository, branch string) (bool, error) {
	apiBase := repo.APIBaseURL
	if apiBase == "" {
		apiBase = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s/%s/pulls?head=%s:%s", strings.TrimRight(apiBase, "/"), repo.Owner, repo.Name, repo.Owner, branch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "token "+repo.AuthToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var prs []map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&prs); err != nil {
			return false, err
		}
		return len(prs) > 0, nil
	}
	return false, fmt.Errorf("GitHub API returned status %d when checking PRs", resp.StatusCode)
}
