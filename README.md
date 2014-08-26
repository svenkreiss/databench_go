# databench_go

> [Go](http://golang.org/) language kernel for [Databench](http://www.svenkreiss.com/databench/). Write your data analysis in Go and visualize and interact with it in the browser.

[![GoDoc](https://godoc.org/github.com/svenkreiss/databench_go/databench?status.png)](https://godoc.org/github.com/svenkreiss/databench_go/databench)
<!--
[![Build Status](https://travis-ci.org/svenkreiss/databench_go.png?branch=master)](https://travis-ci.org/svenkreiss/databench_go)
-->


Install using

    go get github.com/svenkreiss/databench_go/databench

and add `github.com/svenkreiss/databench_go/databench` to your imports.


### Testing / Examples

The `analyses` folder is an example setup and can act as a template for your own Databench analyses with Go. Dependencies for this are installed with

    pip install -r requirements_analyses.txt

![helloworld](docs/helloworld.png)


### Troubleshooting

* `libzmq`: On a Mac, install using `brew install zmq`. To run on linux, make sure you have version 4 of `libzmq`.
