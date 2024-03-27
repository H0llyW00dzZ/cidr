// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package cidr_test

import (
	"fmt"
	"net"

	"github.com/H0llyW00dzZ/cidr"
)

func ExampleIPv4RangeVerify() {
	cidrBlock := "192.168.1.0/24"
	_, ipNet, err := net.ParseCIDR(cidrBlock)
	if err != nil {
		fmt.Printf("Error parsing CIDR block: %v\n", err)
		return
	}

	start, end := cidr.IPv4ToRange(ipNet)
	fmt.Printf("CIDR block %s ranges from %d to %d\n", cidrBlock, start, end)

	// Test the IPv4RangeVerify function
	testIPs := []string{
		"192.168.1.134",   // should be within the CIDR block
		"192.168.2.10",    // should be outside the CIDR block
		"256.256.256.256", // invalid IP address
	}

	for _, ip := range testIPs {
		withinRange, err := cidr.IPv4RangeVerify(cidrBlock, ip)
		if err != nil {
			fmt.Printf("Error verifying IP address '%s': %v\n", ip, err)
		} else {
			fmt.Printf("IP address '%s' is within CIDR block '%s': %t\n", ip, cidrBlock, withinRange)
		}
	}

}

func ExampleIPv4ToRange() {
	_, cidrNet, _ := net.ParseCIDR("10.0.0.0/8")
	start, end := cidr.IPv4ToRange(cidrNet)

	fmt.Printf("Start IP: %d\n", start)
	fmt.Printf("End IP: %d\n", end)

}

func ExampleSingleIPv4ToUint32() {
	ipStr := "192.168.1.1"
	ipNum, err := cidr.SingleIPv4ToUint32(ipStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("IP %s as uint32: %d\n", ipStr, ipNum)
	}

}
