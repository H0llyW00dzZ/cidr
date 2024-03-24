# cidr Range Converter

[![Go Reference](https://pkg.go.dev/badge/github.com/H0llyW00dzZ/cidr.svg)](https://pkg.go.dev/github.com/H0llyW00dzZ/cidr)

The `cidr` package provides a Go utility for converting CIDR (Classless Inter-Domain Routing) blocks into a range of numeric IP addresses. It is designed to work specifically with IPv4 addresses and is useful for network-related operations such as checking if an IP address falls within a CIDR range.

## Features

- Convert CIDR blocks to numeric IP address ranges.
- Handle IPv4 addresses.
- Simple and easy-to-use API.

## Installation

To install the `cidr` package, you need to have Go installed on your machine. Use the following go get command to retrieve the package:

```sh
go get github.com/H0llyW00dzZ/cidr
```

## Usage

Here's a quick example of how to use the `cidr` package:

```go
package main

import (
    "fmt"
    "net"

    "github.com/H0llyW00dzZ/cidr"
)

func main() {
    cidrBlock := "192.168.1.0/24"
    _, ipNet, err := net.ParseCIDR(cidrBlock)
    if err != nil {
        fmt.Printf("Error parsing CIDR block: %v\n", err)
        return
    }

    start, end := cidr.IPv4ToRange(ipNet)
    fmt.Printf("CIDR block %s ranges from %d to %d\n", cidrBlock, start, end)
}
```

Another:

- [The Go Playground](https://go.dev/play/p/Nbh-xA2ecN6)

