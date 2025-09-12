package utils

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"strings"
)

// inferRepoNameFromPath tries to extract repository name from a local path
func inferRepoNameFromPath(path string) string {
	if path == "" {
		return ""
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
func BuildRemoteURL(repo config.Repository, global config.GlobalConfig) string {
	base := strings.TrimSuffix(global.GitBaseURL, "/")
	owner := repo.Owner
	if owner == "" {
		owner = global.Owner
	}
	name := repo.Name
	if name == "" {
		name = inferRepoNameFromPath(repo.Path)
	}

	return fmt.Sprintf("%s/%s/%s.git", base, owner, name)
}
func IsLikelyRemoteBranch(exec executor.Executor, repo config.Repository, base string) bool {
	out, err := exec.RunCommand(repo.Path, "git", "ls-remote", "--heads", "origin", base)
	return err == nil && strings.Contains(out, base)
}
