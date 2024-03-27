// Copyright (c) 2024 H0llyW00dz All rights reserved.
//
// License: BSD 3-Clause License

package cidr

import (
	"net"
	"testing"
)

// TestIPv4ToRange tests the IPv4ToRange function with various CIDR blocks.
func TestIPv4ToRange(t *testing.T) {
	tests := []struct {
		cidr       string
		wantStart  uint32
		wantEnd    uint32
		expectFail bool
	}{
		{"192.168.1.0/24", 0xC0A80100, 0xC0A801FF, false},
		{"10.0.0.0/8", 0x0A000000, 0x0AFFFFFF, false},
		{"172.16.0.0/12", 0xAC100000, 0xAC1FFFFF, false},
		{"0.0.0.0/0", 0x00000000, 0xFFFFFFFF, false},
		{"not a cidr", 0, 0, true}, // Invalid CIDR block should fail.
	}

	for _, tt := range tests {
		_, cidrNet, err := net.ParseCIDR(tt.cidr)
		if err != nil {
			if !tt.expectFail {
				t.Errorf("ParseCIDR(%q) unexpected error: %v", tt.cidr, err)
			}
			continue
		}

		if tt.expectFail {
			t.Errorf("ParseCIDR(%q) expected to fail but didn't", tt.cidr)
			continue
		}

		gotStart, gotEnd := IPv4ToRange(cidrNet)
		if gotStart != tt.wantStart || gotEnd != tt.wantEnd {
			t.Errorf("IPv4ToRange(%q) = %v, %v; want %v, %v", tt.cidr, gotStart, gotEnd, tt.wantStart, tt.wantEnd)
		}
	}
}

// TestIPv4RangeVerify tests the IPv4RangeVerify function with various CIDR blocks and IP addresses.
func TestIPv4RangeVerify(t *testing.T) {
	tests := []struct {
		cidr         string
		ip           string
		expectWithin bool
		expectErr    bool
	}{
		{"192.168.1.0/24", "192.168.1.5", true, false},
		{"192.168.1.0/24", "192.168.1.255", true, false},
		{"192.168.1.0/24", "192.168.2.1", false, false},
		{"10.0.0.0/8", "10.255.255.255", true, false},
		{"10.0.0.0/8", "11.0.0.1", false, false},
		{"0.0.0.0/0", "255.255.255.255", true, false},
		{"not a cidr", "192.168.1.1", false, true},   // Invalid CIDR block should cause an error.
		{"192.168.1.0/24", "not an ip", false, true}, // Invalid IP should cause an error.
	}

	for _, tt := range tests {
		withinRange, err := IPv4RangeVerify(tt.cidr, tt.ip)
		if tt.expectErr {
			if err == nil {
				t.Errorf("IPv4RangeVerify(%q, %q) expected an error, but no error was returned", tt.cidr, tt.ip)
			}
			continue
		}

		if err != nil {
			t.Errorf("IPv4RangeVerify(%q, %q) unexpected error: %v", tt.cidr, tt.ip, err)
			continue
		}

		if withinRange != tt.expectWithin {
			t.Errorf("IPv4RangeVerify(%q, %q) = %v; want %v", tt.cidr, tt.ip, withinRange, tt.expectWithin)
		}
	}
}

// TestSingleIPv4ToUint32 tests the singleIPv4ToUint32 function with various IP addresses.
func TestSingleIPv4ToUint32(t *testing.T) {
	tests := []struct {
		ipStr     string
		want      uint32
		expectErr bool
	}{
		{"192.168.1.1", 0xC0A80101, false},
		{"10.0.0.1", 0x0A000001, false},
		{"255.255.255.255", 0xFFFFFFFF, false},
		{"0.0.0.0", 0x00000000, false},
		{"256.0.0.1", 0, true},  // Invalid IP should cause an error.
		{"-1.0.0.1", 0, true},   // Invalid IP should cause an error.
		{"not.an.ip", 0, true},  // Invalid IP should cause an error.
		{"10.0.0.256", 0, true}, // Invalid IP should cause an error.
	}

	for _, tt := range tests {
		got, err := SingleIPv4ToUint32(tt.ipStr)
		if tt.expectErr {
			if err == nil {
				t.Errorf("singleIPv4ToUint32(%q) expected an error, but no error was returned", tt.ipStr)
			}
			continue
		}

		if err != nil {
			t.Errorf("singleIPv4ToUint32(%q) unexpected error: %v", tt.ipStr, err)
			continue
		}

		if got != tt.want {
			t.Errorf("singleIPv4ToUint32(%q) = %v; want %v", tt.ipStr, got, tt.want)
		}
	}
}
