# databench_go

> [Go](http://golang.org/) language kernel for [Databench](http://www.svenkreiss.com/databench/).

[![GoDoc](https://godoc.org/github.com/svenkreiss/databench_go?status.png)](https://godoc.org/github.com/svenkreiss/databench_go)
[![Build Status](https://travis-ci.org/svenkreiss/databench_go.png?branch=master)](https://travis-ci.org/svenkreiss/databench_go)


Install using

    go get github.com/svenkreiss/databench_go/databench

and add `github.com/svenkreiss/databench_go/databench` to your imports.


### Testing / Examples

The `analyses` folder is an example setup and can act as a template for your own Databench analyses with Go. Dependencies for this are installed with

    pip install -r requirements_analyses.txt


### Troubleshooting

* `libzmq`: On a Mac, install using `brew install zmq`. On Ubuntu, use `sudo apt-get install libzmq3-dev` (somehow libzmq3 installs version 4.0.4 ... weird ... but zmq4 is what you want).
