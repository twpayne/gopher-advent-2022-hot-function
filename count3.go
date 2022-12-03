package main

// A PackedFlow is an IPv4 network flow.
type PackedFlow struct {
	SrcIP   uint32
	DstIP   uint32
	SrcPort uint16
	DstPort uint16
	Proto   byte
}

// CountBadIPs3 returns the number of bad IPs in flows.
func CountBadIPs3(badIPs IPSet, flows []PackedFlow) int {
	totalBadIPs := 0
	for _, flow := range flows {
		if _, ok := badIPs[flow.SrcIP]; ok {
			totalBadIPs++
		}
		if _, ok := badIPs[flow.DstIP]; ok {
			totalBadIPs++
		}
	}
	return totalBadIPs
}
