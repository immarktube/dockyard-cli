package utils

import (
	"github.com/immarktube/dockyard-cli/config"
	"sync"
)

func ForEachRepoConcurrently(repos []config.Repository, fn func(repo config.Repository)) {
	var wg sync.WaitGroup
	wg.Add(len(repos))

	for _, repo := range repos {
		repo := repo // 避免闭包捕获问题
		go func() {
			defer wg.Done()
			fn(repo)
		}()
	}

	wg.Wait()
}
