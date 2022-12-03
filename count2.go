package main

import "net"

type IPSet map[uint32]struct{}

// netIPToUint32 converts an net.IP (which is assumed to be an IPv4 address) to
// a uint32.
func netIPToUint32(netIP net.IP) uint32 {
	return uint32(netIP[0])<<24 + uint32(netIP[1])<<16 + uint32(netIP[2])<<8 + uint32(netIP[3])
}

// newIPSet returns a new set of IPs.
func newIPSet(ips []net.IP) IPSet {
	ipSet := make(IPSet)
	for _, ip := range ips {
		ipSet[netIPToUint32(ip)] = struct{}{}
	}
	return ipSet
}

// CountBadIPs2 returns the number of bad IPs in flows.
func CountBadIPs2(badIPs IPSet, flows []*Flow) int {
	totalBadIPs := 0
	for _, flow := range flows {
		if _, ok := badIPs[netIPToUint32(flow.SrcIP)]; ok {
			totalBadIPs++
		}
		if _, ok := badIPs[netIPToUint32(flow.DstIP)]; ok {
			totalBadIPs++
		}
	}
	return totalBadIPs
}
