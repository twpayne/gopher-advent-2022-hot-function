package main

import "testing"

func BenchmarkCountBadIPs2(b *testing.B) {
	badIPs, flows, _ := getFlowsAndBadIPs(benchNumFlows, benchNumBadIPs, benchHitRate)
	badIPSet, _ := getPackedFlows(badIPs, flows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs2(badIPSet, flows)
	}
	b.StopTimer()
}
