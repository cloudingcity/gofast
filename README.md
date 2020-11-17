# Gofast

[![Test](https://github.com/cloudingcity/gofast/workflows/Test/badge.svg)](https://github.com/cloudingcity/gofast/actions?query=workflow%3ATest)
[![Lint](https://github.com/cloudingcity/gofast/workflows/Lint/badge.svg)](https://github.com/cloudingcity/gofast/actions?query=workflow%3ALint)
[![codecov](https://codecov.io/gh/cloudingcity/gofast/branch/main/graph/badge.svg)](https://codecov.io/gh/cloudingcity/gofast)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/gofast)](https://goreportcard.com/report/github.com/cloudingcity/gofast)

⚡️ Gofast is a HTTP client based on [fasthttp](https://github.com/valyala/fasthttp) with zero memory allocation. 

Automatic struct binding let you focus on entity writing.

> [JSON-to-Go](https://mholt.github.io/json-to-go/) is very useful to generate struct.

## Install

```console
go get -u github.com/cloudingcity/gofast
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/cloudingcity/gofast"
)

type Out struct {
	Hello string `json:"hello"`
}

func main() {
	fast := gofast.New()

	var out Out
	uri := "http://echo.jsontest.com/hello/world"
	if err := fast.Get(uri, &out, nil); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("hello %v", out.Hello)
	// hello world
}
```

## Examples

### Send request with body

```go
type CreateToken struct {
    ID     string `json:"id"`
    Secret string `json:"secret"`
}

type Token struct {
    Token     string `json:"token"`
    ExpiredAt string `json:"expired_at"`
}

fast := gofast.New()

uri := "https://example.com/api/v1/token"
body := CreateToken{
    ID:     "my-id",
    Secret: "my-secret",
}
var token Token
if err := fast.Post(uri, &body, &token, nil); err != nil {
    log.Fatalln(err)
}
fmt.Printf("token: %v, expired_at: %v", token.Token, token.ExpiredAt)
```

### Get with header

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

fast := gofast.New()

var user User
uri := "https://example.com/api/v1/users/100"
h := gofast.Header{fasthttp.HeaderAuthorization: "Bearer My-JWT"}
if err := fast.Get(uri, &user, h); err != nil {
    log.Fatalln(err)
}
fmt.Printf("id: %v, name: %v", user.ID, user.Name)
```

### Customize error handler

Error handler will handle non 200 HTTP status code.

```go
cfg := gofast.Config{
    ErrorHandler: func(resp *fasthttp.Response) error {
        return fmt.Errorf("http code = %d", resp.StatusCode())
    },
}

fast := gofast.New(cfg)
err := fast.Get(uri, nil, nil)
// http code = 400
```

## Benchmarks

```console
$ go test -bench=. -benchmem -benchtime=3s -run=none -cpu 4
BenchmarkClient-4    1000000    3220 ns/op    0 B/op    0 allocs/op
```
