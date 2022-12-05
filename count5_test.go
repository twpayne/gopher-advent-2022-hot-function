package main

import "testing"

func BenchmarkCountBadIPs5(b *testing.B) {
	badIPs, flows, _ := getFlowsAndBadIPs(benchNumFlows, benchNumBadIPs, benchHitRate)
	badIPSet, packedFlows := getPackedFlows(badIPs, flows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs5(badIPSet, packedFlows)
	}
	b.StopTimer()
}
