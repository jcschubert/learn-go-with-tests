# Installation on Windows

## Download and Install Go

* Download the [current version](https://golang.org/dl/) and run the installer.
* Choose a folder to install Go to. (e.g. `C:\go\`)
* The installer should then put `C:\go\bin` into your `PATH`
* Restart open command prompts, and test wether the installation succeeded with `go version`.

## Configure the Go Environment

* All go code lives within a single workspace with a specific file structure.
* This workspace is identified by the environment variable `GOPATH`. By default, go assumes the workspace to be `$HOME/go`.
* Create a workspace folder (e.g. `D:\go`).
* Add the Environment Variable `GOPATH` with the value `D:\go`, and add `%GOPATH%\bin` to the end of `PATH`.
* Do an `echo $GOPATH` and `echo $PATH` to verify your settings.

## Create the Go Environment Folder Structure

```bash
mkdir -p $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

* verify with `cd $GOPATH` and `ls`
* Now, `go get`, the built-in package manager, should work correctly.

## Install an Editor

* No preference? Use [VS Code](https://code.visualstudio.com/), it has *excellent* go support.
* Be sure to install the go extension. You might also be prompted for additional helpers, install them, too.

## Install the Go Debugger

* `go get -u github.com/go-delve/delve/cmd/dlv`.
* Check `$GOPATH/bin` afterwards.

## Install the Go Linter

* `go get -u github.com/golangci/golangci-lint/cmd/golangci-lint`.
* Check `$GOPATH/bin` afterwards.
