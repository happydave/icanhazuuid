package server

import (
	"strings"
	"sync/atomic"
)

var TotalRequests uint64
var AddressRequestCounts = make(map[string]uint64)

func addRequest(ip string) {
	atomic.AddUint64(&TotalRequests, 1)

	ip = strings.Split(ip, ":")[0]

	val, ok := AddressRequestCounts[ip]
	if !ok {
		AddressRequestCounts[ip] = 0
	}

	AddressRequestCounts[ip] = val + 1
}

func getTotalRequestCount() uint64 {
	return TotalRequests
}

func getRequestCountForIP(ip string) (uint64, string) {
	ip = strings.Split(ip, ":")[0]

	val, ok := AddressRequestCounts[ip]
	if !ok {
		val = 0
	}

	return val, ip
}
