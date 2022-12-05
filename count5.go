package main

import (
	"runtime"
	"sync"
)

// CountBadIPs5 returns the number of bad IPs in flows.
func CountBadIPs5(badIPs IPSet, flows []PackedFlow) int {
	var wg sync.WaitGroup
	numGoroutines := runtime.NumCPU()
	chunkSize := len(flows) / numGoroutines
	badIPsByGoroutine := make([]int, numGoroutines)
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()

			// Goroutine 0 gets the first 1/Nth chunk
			// Goroutine 1 gets the second 1/Nth chunk
			// Goroutine 2 gets the third 1/Nth chunk
			// etc.
			chunkStart := i * chunkSize
			chunkEnd := (i + 1) * chunkSize
			if chunkEnd > len(flows) {
				chunkEnd = len(flows)
			}

			for j := chunkStart; j < chunkEnd; j++ {
				flow := flows[j]
				if _, ok := badIPs[flow.SrcIP]; ok {
					badIPsByGoroutine[i]++
				}
				if _, ok := badIPs[flow.DstIP]; ok {
					badIPsByGoroutine[i]++
				}
			}
		}(i)
	}

	wg.Wait()

	totalBadIPs := 0
	for i := 0; i < numGoroutines; i++ {
		totalBadIPs += badIPsByGoroutine[i]
	}
	return totalBadIPs
}
