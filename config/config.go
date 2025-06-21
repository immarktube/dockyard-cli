package config

import (
	"fmt"
	"os"
	_ "strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Repository struct {
	Path       string `yaml:"path"`
	Hook       *Hook  `yaml:"hook,omitempty"`
	URL        string `yaml:"url,omitempty"`
	Owner      string `yaml:"owner,omitempty"`
	Name       string `yaml:"name,omitempty"`
	AuthToken  string `yaml:"authToken,omitempty"`
	APIBaseURL string `yaml:"apiBaseURL,omitempty"`
	Branch     string `yaml:"branch,omitempty"`
}

type GlobalConfig struct {
	Owner       string `yaml:"owner,omitempty"`
	AuthToken   string `yaml:"authToken,omitempty"`
	APIBaseURL  string `yaml:"apiBaseURL,omitempty"`
	Concurrency int    `yaml:"concurrency,omitempty"`
}

type Hook struct {
	Pre  string `yaml:"pre"`
	Post string `yaml:"post"`
}

type Config struct {
	Repositories []Repository      `yaml:"repositories"`
	Global       GlobalConfig      `yaml:"global,omitempty"`
	Hooks        Hook              `yaml:"hook"`
	Env          map[string]string // load environment variable from .env
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	env, err := godotenv.Read()
	if err == nil {
		cfg.Env = env
	} else {
		cfg.Env = make(map[string]string)
	}

	yamlFile, err := os.ReadFile(".dockyard.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read .dockyard.yaml: %w", err)
	}
	if err := yaml.Unmarshal(yamlFile, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse .dockyard.yaml: %w", err)
	}
	for k, v := range cfg.Env {
		if _, exists := os.LookupEnv(k); !exists {
			os.Setenv(k, v)
		}
	}

	// Expand environment variables in global config
	cfg.Global.AuthToken = os.ExpandEnv(cfg.Global.AuthToken)
	cfg.Global.APIBaseURL = os.ExpandEnv(cfg.Global.APIBaseURL)
	cfg.Global.Owner = os.ExpandEnv(cfg.Global.Owner)

	for i := range cfg.Repositories {
		repo := &cfg.Repositories[i]

		repo.Path = os.ExpandEnv(repo.Path)
		repo.URL = os.ExpandEnv(repo.URL)
		repo.Owner = os.ExpandEnv(repo.Owner)
		repo.Name = os.ExpandEnv(repo.Name)
		repo.AuthToken = os.ExpandEnv(repo.AuthToken)
		repo.APIBaseURL = os.ExpandEnv(repo.APIBaseURL)
		repo.Branch = os.ExpandEnv(repo.Branch)

		if repo.Owner == "" {
			repo.Owner = cfg.Global.Owner
		}
		if repo.AuthToken == "" {
			repo.AuthToken = cfg.Global.AuthToken
		}
		if repo.APIBaseURL == "" {
			repo.APIBaseURL = cfg.Global.APIBaseURL
		}
		if repo.Name == "" {
			repo.Name = inferRepoNameFromPath(repo.Path)
		}
	}

	return cfg, nil
}

func GetHooksForRepo(cfg *Config, repo Repository) Hook {
	if repo.Hook != nil && (repo.Hook.Pre != "" || repo.Hook.Post != "") {
		return *repo.Hook
	}
	return cfg.Hooks
}

// inferRepoNameFromPath tries to extract repository name from local path
func inferRepoNameFromPath(path string) string {
	if path == "" {
		return ""
	}
	// Strip trailing slash if present
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	// Return last segment
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
