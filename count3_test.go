package main

import (
	"net"
	"testing"
)

var (
	benchBadIPSet    IPSet
	benchPackedFlows []PackedFlow
)

func BenchmarkCountBadIPs3(b *testing.B) {
	badIPs, flows, _ := getFlowsAndBadIPs(benchNumFlows, benchNumBadIPs, benchHitRate)
	badIPSet, packedFlows := getPackedFlows(badIPs, flows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs3(badIPSet, packedFlows)
	}
	b.StopTimer()
}

func getPackedFlows(badIPs []net.IP, flows []*Flow) (IPSet, []PackedFlow) {
	if benchBadIPSet == nil {
		benchBadIPSet = newIPSet(badIPs)
		benchPackedFlows = packFlows(flows)
	}
	return benchBadIPSet, benchPackedFlows
}

func packFlows(flows []*Flow) []PackedFlow {
	packedFlows := make([]PackedFlow, 0, len(flows))
	for _, flow := range flows {
		packedFlow := PackedFlow{
			SrcIP:   netIPToUint32(flow.SrcIP),
			SrcPort: flow.SrcPort,
			DstIP:   netIPToUint32(flow.DstIP),
			DstPort: flow.DstPort,
			Proto:   flow.Proto,
		}
		packedFlows = append(packedFlows, packedFlow)
	}
	return packedFlows
}
