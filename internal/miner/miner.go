package miner

import (
	"runtime"
	"sync"
)

// Mine starts multiple workers and returns the first successful result.
func Mine(
	difficulty int,
	hashFunc func(uint64) (string, error),
) (Result, error) {

	workers := runtime.NumCPU()

	results := make(chan Result, 1)
	stop := make(chan struct{})

	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {

		wg.Add(1)

		go func(start uint64) {
			defer wg.Done()

			Worker(
				start,
				uint64(workers),
				difficulty,
				hashFunc,
				results,
				stop,
			)
		}(uint64(i))
	}

	result := <-results

	once.Do(func() {
		close(stop)
	})

	wg.Wait()

	return result, nil
}
