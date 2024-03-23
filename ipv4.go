// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package cidr

import (
	"net"
)

// IPv4ToRange converts a CIDR block to a range of numeric IP addresses.
// It takes a [*net.IPNet] which represents the CIDR block and returns the start
// and end IP addresses as uint32 integers. This function is designed to work
// with IPv4 addresses only.
//
// Parameters:
//
//	cidr: A pointer to a net.IPNet struct representing the CIDR block.
//
// Returns:
//
//	start: The first IP address in the range as a uint32.
//	end: The last IP address in the range as a uint32.
//
// The function assumes that the provided CIDR block is valid and that [*net.IP.To4]
// is an IPv4 address. If cidr or cidr.IP is nil, the behavior of the function
// is undefined.
func IPv4ToRange(cidr *net.IPNet) (uint32, uint32) {
	ip := cidr.IP.To4() // Ensure we are dealing with an IPv4 address.
	if ip == nil {
		// This should not happen if the preconditions are met.
		// If it does, returning 0,0 is a safe fallback.
		return 0, 0
	}

	var start uint32
	for _, bytePart := range ip {
		start = (start << 8) | uint32(bytePart) // Convert IP to a 32-bit number.
	}

	// Calculate the end of the range based on the CIDR mask.
	ones, bits := cidr.Mask.Size()           // Get mask as the number of leading ones and total bits.
	mask := uint32((1 << (bits - ones)) - 1) // Ensure mask is of type uint32.
	end := start | mask                      // Set the remaining bits to 1 to get the last address in the range.

	return start, end
}
