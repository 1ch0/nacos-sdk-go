# nacos-client
[![build workflow](https://github.com/1ch0/nacos-client/actions/workflows/badge.yml/badge.svg)](https://github.com/1ch0/nacos-client/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/1ch0/nacos-client)](https://pkg.go.dev/github.com/1ch0/nacos-client?tab=doc)

> nacos-client is light weight nacos client.

## Documentation

- [English](./README.md)
- [简体中文](./README_zh.md)

## Resources

- [Reference](https://pkg.go.dev/github.com/1ch0/nacos-client)
- [Examples](https://pkg.go.dev/github.com/1ch0/nacos-client#pkg-examples)

## Features

- [Configuration Management](https://nacos.io/en-us/docs/open-api.html)
- [Service Discover](https://nacos.io/en-us/docs/open-api.html)
- [Namespace](https://nacos.io/en-us/docs/open-api.html)
- Permisstion Management
  - User Management
  - Role Management
  - Permisstions Management


## Installation

nacos-client supports 2 last Go versions and requires a Go version with [modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go module:

```shell
go mod init github.com/my/repo
```

Then install nacos-client:

```shell
go get github.com/1ch0/nacos-client
```

## Quickstart

### Connecting via a nacos url

```go
import (
	nacos "github.com/1ch0/nacos-client"
)

func ExampleClient() {
	client := nacos.New(
		&nacos.Config{
			Addr:     "http://locahost:8848",
			Username: "nacos",
			Password: "nacos",
		})

	if err := client.Health(); err != nil {
		panic(err)
	}

	if err := client.Login(); err != nil {
		panic(err)
	}
}
```

## Look and feel

Some corner cases:

```go

```

## Run the test

