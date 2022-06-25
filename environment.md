# Editor Setup

## Installation

The latest version of the tools can be found at the downloads page on the Go website. Instructions for downloading and installing Go
[Installing Go](https://golang.org/dl)

If you are a Mac developer, you can install Go using [Homebrew](https://brew.sh) with the command  `brew install go`


Windows Developers who use [Chocolatey](https://chocolatey.org) can install Go with the command `choco install golang`

A document that summarizes commonly used editor plugins and IDEs with Go support
[Editor plugins and IDEs](https://go.dev/doc/editors.html)

Note: Go programs compile to a single binary and do not require any additional software to be installed in order to run them. Install the Go development tools only on computers that build Go programs.

You can validate that your environment is set up correctly by opening up a terminal or command prompt and typing:

 `$ go version`

If everythink is set up correctly, you should see something like this printed:

`go version go1.17.8 linux/amd64`

If you get an error instead of the version message, it's likely that you don't have go in your executable path, or you have another program named go in your path. 

## Formatting


The Go development tools include a command, `go fmt`, which automatically reformats your code to match the standard format. It does things like fixing up the whitespace for indenting, lining up fields in a struct, and making shure here is proper spacing around operators.

## Linting

A tool called golint tries to ensure that your code follows style guidelines. Some of the changes it suggests include properly naming variables, formatting error messages, and placing comments on public methods and types.

Install golint with the following command:

`go install golang.org/x/lint/golint@latest`

And run it with:

`golint ./...`

This runs golint over your entire project.

## Vetting

Sometimes the code is syntactically valid, but there are mistakes that are not what you meant to do. This includes things like passing the wrong number of parameters to formatting methods or assigning values to variables that are never used. The go tool includes a command called `go vet` to detect these kinds of errors. Run go vet on your code with the command:

`go vet ./...`


## Makefiles

Go developers have adopted `make` to ensure repeatable, automated builds that can be run by anyone, anywhere, at any time.

Here's a sample Makefile:

```
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build *.go
.PHONY:vet
```

[Back](README.md){: .btn}