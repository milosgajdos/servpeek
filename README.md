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

Currently there is no single program to run as it stands. Instead the project provides API to build your own go binaries that can be used to test your infrastrucure - this way there is no need to have any runtime available on the tested infrastructure, just your binary. You can obviously create your own `_test.go` files for the stuff you want to test, but you will need to have go tool chain available on the tested server. Building binaries is recommended and preferrable option.

## Example

Get the package:
```
$ go get github.com/milosgajdos83/servpeek
```

You can find an example program in the [examples](https://github.com/milosgajdos83/servpeek/tree/master/examples) directory.

## TODO

- A LOT more resources
- Better logging and error statements
