package main

import (
	"math/rand"
	"net"
	"testing"
)

const (
	benchNumFlows  = 1000000
	benchNumBadIPs = 10000
	benchHitRate   = 0.01
)

var (
	benchBadIPs []net.IP
	benchFlows  []*Flow
	benchHits   int
)

func BenchmarkCountBadIPs1(b *testing.B) {
	badIPs, flows, _ := getFlowsAndBadIPs(benchNumFlows, benchNumBadIPs, benchHitRate)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = CountBadIPs1(badIPs, flows)
	}
	b.StopTimer()
}

func getFlowsAndBadIPs(numBadIPs, numFlows int, hitRate float64) ([]net.IP, []*Flow, int) {
	if benchBadIPs == nil {
		benchBadIPs, benchFlows, benchHits = generateRandomBadIPsAndFlows(benchNumBadIPs, benchNumFlows, benchHitRate)
	}
	return benchBadIPs, benchFlows, benchHits
}

func generateRandomBadIPsAndFlows(numBadIPs, numFlows int, hitRate float64) ([]net.IP, []*Flow, int) {
	r := rand.New(rand.NewSource(0))
	badIPs := make([]net.IP, 0, numBadIPs)
	for i := 0; i < numBadIPs; i++ {
		badIPs = append(badIPs, randomIP(r))
	}
	flows := make([]*Flow, 0, numFlows)
	hits := 0
	for i := 0; i < numFlows; i++ {
		var flow Flow
		if r.Float64() < hitRate {
			flow.SrcIP = badIPs[r.Intn(len(badIPs))]
			hits++
		} else {
			flow.SrcIP = randomIP(r)
		}
		if r.Float64() < hitRate {
			flow.DstIP = badIPs[r.Intn(len(badIPs))]
			hits++
		} else {
			flow.DstIP = randomIP(r)
		}
		flows = append(flows, &flow)
	}
	return badIPs, flows, hits
}

func randomIP(r *rand.Rand) net.IP {
	bytes := make([]byte, 4)
	if n, err := r.Read(bytes); n != len(bytes) || err != nil {
		panic("math/rand.Source.Read failed")
	}
	return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3])
}
