package utils

import (
	"github.com/immarktube/dockyard-cli/config"
	"sync"
)

func ForEachRepoConcurrently(repos []config.Repository, fn func(repo config.Repository), maxConcurrency int) {
	if maxConcurrency <= 0 {
		maxConcurrency = 5
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	for _, repo := range repos {
		wg.Add(1)
		repo := repo // capture the current repo in the loop
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}        // acquire a token
			defer func() { <-semaphore }() // release the token
			fn(repo)
		}()
	}

	wg.Wait()
}

func GetConcurrency(cmdConcurrency int, cfg *config.Config) int {
	if cmdConcurrency > 0 {
		return cmdConcurrency
	}
	if cfg != nil && cfg.Concurrency > 0 {
		return cfg.Concurrency
	}
	return 5
}
