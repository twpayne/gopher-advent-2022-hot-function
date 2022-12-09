package main

import (
	"flag"
	"math"
	"runtime"
)

var overProvisionGoroutines = flag.Float64("over-provision-goroutines", 1, "Over-provision goroutines")

// CountBadIPs7 returns the number of bad IPs in flows.
func CountBadIPs7(badIPs IPSet, flows []PackedFlow) int {
	numGoroutines := int(math.Round(float64(runtime.NumCPU()) * *overProvisionGoroutines))
	chunkSize := (len(flows) + numGoroutines - 1) / numGoroutines
	countCh := make(chan int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			chunkStart := i * chunkSize
			chunkEnd := (i + 1) * chunkSize
			if chunkEnd > len(flows) {
				chunkEnd = len(flows)
			}

			count := 0
			for _, flow := range flows[chunkStart:chunkEnd] {
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
