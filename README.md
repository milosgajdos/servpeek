# servpeek

Introspective peek into your server guts

## Motivation

Currently the most widely used tools for testing your infrastructure are [serverspec](https://github.com/mizzy/serverspec) and [infrataster](https://github.com/ryotarai/infrataster).

Hoever these being written in [ruby](https://www.ruby-lang.org/en/) they carry a decent amount of runtime an gem dependencies. Writing a simple tool in [Go](http://golang.org/) programming language would help to address these and would further simplify the ditribution and speed of deployment.
