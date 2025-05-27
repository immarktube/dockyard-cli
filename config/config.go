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
	Owner      string `yaml:"owner,omitempty"`
	AuthToken  string `yaml:"authToken,omitempty"`
	APIBaseURL string `yaml:"apiBaseURL,omitempty"`
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
	Concurrency  int               `yaml:"concurrency"`
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

	// Apply global defaults
	for i := range cfg.Repositories {
		repo := &cfg.Repositories[i]
		if repo.Owner == "" {
			repo.Owner = cfg.Global.Owner
		}
		if repo.AuthToken == "" {
			repo.AuthToken = cfg.Global.AuthToken
		}
		if repo.APIBaseURL == "" {
			repo.APIBaseURL = cfg.Global.APIBaseURL
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
