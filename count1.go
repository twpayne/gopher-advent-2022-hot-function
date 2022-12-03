package main

import "net"

// A Flow is an IPv4 network flow.
type Flow struct {
	SrcIP   net.IP
	SrcPort uint16
	DstIP   net.IP
	DstPort uint16
	Proto   byte
}

// isBadIP returns if ip is in badIPs.
func isBadIP(ip net.IP, badIPs []net.IP) bool {
	for _, badIP := range badIPs {
		if ip.Equal(badIP) { // uses bytes.Equal
			return true
		}
	}
	return false
}

// CountBadIPs1 returns the number of bad IPs in flows.
func CountBadIPs1(badIPs []net.IP, flows []*Flow) int {
	totalBadIPs := 0
	for _, flow := range flows {
		if isBadIP(flow.SrcIP, badIPs) {
			totalBadIPs++
		}
	}
	for _, flow := range flows {
		if isBadIP(flow.DstIP, badIPs) {
			totalBadIPs++
		}
	}
	return totalBadIPs
}
