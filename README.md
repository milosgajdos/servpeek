# servpeek

[![GoDoc](https://godoc.org/github.com/milosgajdos/servpeek?status.svg)](https://godoc.org/github.com/milosgajdos/servpeek)
[![Travis CI](https://travis-ci.org/milosgajdos/servpeek.svg?branch=master)](https://travis-ci.org/milosgajdos/servpeek)

Introspective peek into your infrastructure guts

## Motivation

Currently the most widely used tools for testing your infrastructure are [serverspec](https://github.com/mizzy/serverspec) and [infrataster](https://github.com/ryotarai/infrataster).

However these tools being written in [ruby](https://www.ruby-lang.org/en/) carry a decent amount of runtime an gem dependencies. Writing a simple tool in [Go](http://golang.org/) programming language would help to address runtime dependency "issue" and would further simplify the ditribution and speed of deployment of infrastructure.

## Current state

Current project design might be a bit of an overkill mostly because of my Go incompetence or because the projects tries to create abstractions on top of already existing abstractions (package managers, commands etc.). Turns out modelling this is not that easy :-/

Project currently provides support for various package managers: `apt`, `yum`, `apk`, `gem` and `pip` as well as file and service inspections. More resources will be added later as the project moves on [if].

# Usage

The project does not offer a single program to run as it stands right now. Instead a simple API is provided that allows you to build your own `go` binaries that can then be used to test your infrastrucure. Statically linked binary alleviates a necessity to have any language runtime available on the tested infrastructure servers. Obviously, you have an option to create your own `_test.go` files that can be run using `go test`, but like already said you will need to have the `go` tool chain available on the tested server. The choice is really up to you, but my recommendation is to build and shipt the binaries.

## Example

Get the package:
```
$ go get github.com/milosgajdos/servpeek
```

Write a go program that does some testing:
```go
package main

import (
	"log"

	"github.com/milosgajdos/servpeek/file"
	"github.com/milosgajdos/servpeek/pkg"
)

func main() {
	// Test if a gem package is installed
	gemPkg, err := pkg.NewPackage("gem", "bundler", "1.10.6")
	if err != nil {
		log.Fatal(err)
	}

	if err := pkg.IsInstalled(gemPkg); err != nil {
		log.Fatal(err)
	}

	// Test if /etc/hosts is a regular file
	f := file.NewFile("/etc/hosts")
	if err := file.IsRegular(f); err != nil {
		log.Fatal(err)
	}
}
```

You can find a more elaborate example program in the [examples](https://github.com/milosgajdos/servpeek/tree/master/examples) directory.

## TODO

- A LOT more resources
- Better logging and error statements
