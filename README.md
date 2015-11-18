# servpeek

[![GoDoc](https://godoc.org/github.com/milosgajdos83/servpeek?status.svg)](https://godoc.org/github.com/milosgajdos83/servpeek)

Introspective peek into your server guts

## Motivation

Currently the most widely used tools for testing your infrastructure are [serverspec](https://github.com/mizzy/serverspec) and [infrataster](https://github.com/ryotarai/infrataster).

Hoever these being written in [ruby](https://www.ruby-lang.org/en/) they carry a decent amount of runtime an gem dependencies. Writing a simple tool in [Go](http://golang.org/) programming language would help to address these and would further simplify the ditribution and speed of deployment.

## Current state

Current design might be a bit of a overkill mostly because of my Golang incompetence or because the projects tries to create abstractions on top of already existing abstractions (package managers, commands etc.). Turns out modelling this is not very easy :-/

Project currently provides support for various package managers: `apt`, `yum`, `apk`, `gem` and `pip` as well as file inspections. More resources will be added later as the project moves on [if].

# Usage

Currently there is no single program to run as it stands. Instead you can simply invoke the familiar `go test` on any of the examples in `examples` subdirectory. You can obviously create your own `_test.go` files for the stuff you want to test.

## Example

Get the package:
```
$ go get github.com/milosgajdos83/servpeek
```

Example test file could look like this:

```go
package examples

import (
	"testing"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/resource/pkg"
)

func TestGemPackage(t *testing.T) {
	testPkg := resource.Pkg{
		Name:    "bundler",
		Version: "1.10.6",
		Type:    "gem",
	}

	if err := pkg.IsInstalledVersion(testPkg); err !=nil {
		t.Errorf("Error: %s", err)
	}
}
```

Running the tests:
```bash
$ go test gem_package_test.go
ok  	command-line-arguments	0.330s
```

TOD:
- A LOT more resources
- Better logging and error statements
- redesign if possible :-)
