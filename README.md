<pre>
   ______      ______________  ____ 
  / ____/___  / ____/  _/ __ \/ __ \
 / / __/ __ \/ /    / // / / / /_/ /
/ /_/ / /_/ / /____/ // /_/ / _, _/    Range Converter
\____/\____/\____/___/_____/_/ |_|  
                                    
Copyright (©️) 2024 @H0llyW00dzZ All rights reserved.
</pre>

[![Go Version](https://img.shields.io/badge/1.22.1-gray?style=flat&logo=go&logoWidth=15)](https://github.com/H0llyW00dzZ/cidr/blob/master/go.mod#L3)
[![Go Reference](https://pkg.go.dev/badge/github.com/H0llyW00dzZ/cidr.svg)](https://pkg.go.dev/github.com/H0llyW00dzZ/cidr)
[![Go Report Card](https://goreportcard.com/badge/github.com/H0llyW00dzZ/cidr)](https://goreportcard.com/report/github.com/H0llyW00dzZ/cidr)

The `cidr` package provides a Go utility for converting CIDR (Classless Inter-Domain Routing) blocks into a range of numeric IP addresses. It is designed to work specifically with IPv4 addresses and is useful for network-related operations such as checking if an IP address falls within a CIDR range.

## Features

- Convert CIDR blocks to numeric IP address ranges.
- Handle IPv4 addresses.
- Simple and easy-to-use API.

## Installation

> [!NOTE]
> This requires `go1.22.1+`. The reason it's not supported on older versions (e.g, `go1.21` or `lower`)
> is because `go1.22+` supports range over integers. If you're advanced in Go, you can easily perform mass checking using multiple goroutines or a single goroutine.

To install the `cidr` package, you need to have Go installed on your machine. Use the following go get command to retrieve the package:

```sh
go get github.com/H0llyW00dzZ/cidr
```
