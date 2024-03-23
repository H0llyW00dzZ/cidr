// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

// Package cidr provides utilities to work with CIDR (Classless Inter-Domain Routing)
// blocks, specifically for converting them to a range of IP addresses.
//
// The IPv4ToRange function is the core utility of this package, which converts a CIDR block
// into a numeric range of IP addresses represented as uint32 integers. This function is
// particularly useful when working with IPv4 addresses and needing to perform operations
// such as checking if an IP address falls within a certain range.
//
// Example Usage:
//
// The following example demonstrates how to use the IPv4ToRange function to convert a CIDR block
// into a numeric IP range:
//
//	package main
//
//	import (
//	    "fmt"
//	    "net"
//
//	    "github.com/H0llyW00dzZ/cidr"
//	)
//
//	func main() {
//	    // Define a CIDR block as a string.
//	    cidrBlock := "192.168.1.0/24"
//
//	    // Parse the string to a net.IPNet struct.
//	    _, cidr, err := net.ParseCIDR(cidrBlock)
//	    if err != nil {
//	        fmt.Printf("Error parsing CIDR block: %v\n", err)
//	        return
//	    }
//
//	    // Call the [*cidr.IPv4ToRange] function with the parsed CIDR block.
//	    start, end := cidr.IPv4ToRange(cidr)
//
//	    // Print the result.
//	    fmt.Printf("CIDR block %s ranges from %d to %d\n", cidrBlock, start, end)
//	}
//
// Note that the above example assumes that the cidr package is imported correctly, and it will
// print the start and end IP addresses as uint32 integers for the given CIDR block.
package cidr
