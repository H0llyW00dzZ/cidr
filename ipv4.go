// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package cidr

import (
	"net"
	"strconv"
	"strings"
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

// IPv4RangeVerify checks if the given IP address falls within the given CIDR range.
// It returns true if the IP is within the range, false otherwise.
//
// Parameters:
//
//	cidrBlock: A string representing the CIDR block.
//	ipStr: A string representing the IP address to verify.
//
// Returns:
//
//	withinRange: A boolean indicating whether the IP is within the CIDR range.
//	err: An error if there is an issue parsing the CIDR block or IP address.
func IPv4RangeVerify(cidrBlock, ipStr string) (bool, error) {
	_, cidrNet, err := net.ParseCIDR(cidrBlock)
	if err != nil {
		return false, err
	}

	start, end := IPv4ToRange(cidrNet)

	ipNum, err := singleIPv4ToUint32(ipStr)
	if err != nil {
		return false, err
	}

	return ipNum >= start && ipNum <= end, nil
}

// singleIPv4ToUint32 converts an IPv4 address string to its uint32 representation.
//
// This function parses an IPv4 address in dotted-decimal notation ("192.168.1.1")
// and converts it to a uint32 integer. Each octet of the IPv4 address is assumed
// to be in the range of 0 to 255. The function returns an error if the string
// is not a valid representation of an IPv4 address.
//
// Parameters:
//
// ipStr: A string representing the IPv4 address.
//
// Returns:
//
// ipNum: The uint32 representation of the IPv4 address.
// err: An error if the string is not a valid IPv4 address.
func singleIPv4ToUint32(ipStr string) (uint32, error) {
	var ipNum uint32
	bytes := strings.Split(ipStr, ".")
	if len(bytes) != 4 {
		return 0, net.InvalidAddrError("Invalid IP address format")
	}
	for _, b := range bytes {
		p, err := strconv.ParseUint(b, 10, 8)
		if err != nil {
			return 0, err
		}
		ipNum = (ipNum << 8) + uint32(p)
	}
	return ipNum, nil
}
