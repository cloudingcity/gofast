# Gofast

[![Test](https://github.com/cloudingcity/gofast/workflows/Test/badge.svg)](https://github.com/cloudingcity/gofast/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gofast/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gofast/actions?query=workflow%3ALint)
[![codecov](https://codecov.io/gh/cloudingcity/gofast/branch/main/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gofast)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gofast)](https://goreportcard.com/report/github.com/cloudingcity/gofast)

⚡️ Gofast is a HTTP client based on [fasthttp](https://github.com/valyala/fasthttp) with zero memory allocation. 

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/cloudingcity/gofast"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fast := gofast.New()

	var user User
	rawURL := "https://example.com/api/v1/users/100"
	if err := fast.Get(rawURL, &user, nil); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("id: %v, name: %v", user.ID, user.Name)
}
```

## Benchmarks

```console
$ go test -bench=. -benchmem -benchtime=3s -run=none -cpu 4
BenchmarkClient-4    1000000    3220 ns/op    0 B/op    0 allocs/op
```

## Install

```console
go get -u github.com/cloudingcity/gofast
```
