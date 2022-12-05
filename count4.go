package main

import (
	"runtime"
	"sync"
)

// CountBadIPs4 returns the number of bad IPs in flows.
func CountBadIPs4(badIPs IPSet, flows []PackedFlow) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	totalBadIPs := 0
	numGoroutines := runtime.NumCPU()
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			// Goroutine 0 processes 0, numGoroutines, 2*numGoroutines
			// Goroutine 1 processes 1, numGoroutines+1, 2*numGoroutines+1
			// Goroutine 2 processes 2, numGoroutines+2, 2*numGoroutines+2
			// etc.
			for j := i; j < len(flows); j += numGoroutines {
				flow := flows[j]
				if _, ok := badIPs[flow.SrcIP]; ok {
					mu.Lock()
					totalBadIPs++
					mu.Unlock()
				}
				if _, ok := badIPs[flow.DstIP]; ok {
					mu.Lock()
					totalBadIPs++
					mu.Unlock()
				}
			}
		}(i)
	}

	wg.Wait()

	return totalBadIPs
}
