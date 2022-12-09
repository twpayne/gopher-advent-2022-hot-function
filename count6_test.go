package main

import "testing"

func BenchmarkCountBadIPs6(b *testing.B) {
	badIPs, flows, _ := getBadIPsAndFlows()
	badIPSet, packedFlows := getPackedFlows(badIPs, flows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs6(badIPSet, packedFlows)
	}
	b.StopTimer()
}
