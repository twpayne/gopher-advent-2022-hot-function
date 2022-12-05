package main

import (
	"runtime"
)

// CountBadIPs6 returns the number of bad IPs in flows.
func CountBadIPs6(badIPs IPSet, flows []PackedFlow) int {
	numGoroutines := runtime.NumCPU()
	chunkSize := len(flows) / numGoroutines
	countCh := make(chan int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			chunkStart := i * chunkSize
			chunkEnd := (i + 1) * chunkSize
			if chunkEnd > len(flows) {
				chunkEnd = len(flows)
			}

			count := 0
			for j := chunkStart; j < chunkEnd; j++ {
				flow := flows[j]
				if _, ok := badIPs[flow.SrcIP]; ok {
					count++
				}
				if _, ok := badIPs[flow.DstIP]; ok {
					count++
				}
			}

			countCh <- count
		}(i)
	}

	totalBadIPs := 0
	for i := 0; i < numGoroutines; i++ {
		totalBadIPs += <-countCh
	}
	close(countCh)
	return totalBadIPs
}
