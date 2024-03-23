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
