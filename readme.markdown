# Go String Utilities

The following repo is a collection of string functions I have created over the
years, and slowly moving them to a single package, helping me and others to enjoy
them, and stop inventing the wheel every project.

[![Go Report Card](https://goreportcard.com/badge/github.com/ik5/gostrutils)](https://goreportcard.com/report/github.com/ik5/gostrutils)
[![GolangCI](https://golangci.com/badges/github.com/ik5/gostrutils.svg)](https://golangci.com/r/github.com/ik5/gostrutil)
[![Build Status](https://travis-ci.org/ik5/gostrutils.svg?branch=master)](https://travis-ci.org/ik5/gostrutils)
[![Coverage Status](https://coveralls.io/repos/github/ik5/gostrutils/badge.svg?branch=master)](https://coveralls.io/github/ik5/gostrutils?branch=master)
[![GoDoc](https://godoc.org/github.com/ik5/gostrutils?status.svg)](https://godoc.org/github.com/ik5/gostrutils)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)


Package gostrutils contains string utilities that are missing from the main
strings package, as well as other encoding packages arrives by go.

The implementation of the package is set by files based on a subject that belong
to.


## Basics

The basic logic of Go, is to use bytes as a non numeric holder.
The basics of the following package is to gain the ability to hold more support
for string based on functions that are missing, while remembering that Go's strings
are UTF-8.

The project itself built with files that holds the subject of what they are doing.

The helpers.go file, holds functions that are not string related functions, but
help creating that string support.

## Test coverage

The aim of the following library is to have close to 100% of unit test coverage,
and also examples for all existed functions.


## Dependencies

The package built to support golang standard library, for minimizing dependencies.
There is no use of any 3rd party packages when using this package, and every
contribution to the project must also take that in consideration.

# License

The Following project released under MIT.
