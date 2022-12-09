package main

import "testing"

func BenchmarkCountBadIPs7(b *testing.B) {
	badIPs, flows, _ := getBadIPsAndFlows()
	badIPSet, packedFlows := getPackedFlows(badIPs, flows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs7(badIPSet, packedFlows)
	}
	b.StopTimer()
}
