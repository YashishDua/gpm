# gpm (Go Package Manager)

gpm provides support to existing Go commands to accomplish project tasks.

# Overview

gpm provides commands:
* To create boilerplate project directory
* To create module inside/outside GOPATH
* To update Go version
* To build project either using vendor or modules
* To generate vendor folder

# Commands

## Help

```bash
$ gpm 

Usage:
gpm [command]

Available Commands:
  build       Build the project
  create      Creates directory structure
  help        Help about any command
  init        Initializes the project
  mod         Creates modules file
  update      Updates Go version
  vendor      Creates vendor using modules
  version     Print the version number of gpm

Flags:
  -h, --help   help for gpm

Use "gpm [command] --help" for more information about a command.
```

## Initialize
Make sure to initialize project with gpm to use any command.

```bash
$ gpm init

# Initializing gpm...
gpm: Initialized
```

## Create
This commands creates a boilerplate project structure.

```bash
$ gpm create

# Setting up project structure...
gpm: Creating cmd directory
gpm: Creating internal directory
gpm: Creating pkg directory
gpm: Creating scripts directory
gpm: Creating api directory
gpm: Creating test directory
gpm: Create successful
```

## Build
This command builds project using vendor or modules as specified and also takes care of whether the project is inside or outside GOPATH.

- To build using modules
```bash
$ gpm build -m
```

- To build using vendor
```bash
$ gpm build -v
```

## Vendor
This commands help to create vendor using modules.

```bash
$ gpm vendor

# Creating vendor...
gpm: using modules to build vendor
gpm: Vendor created
```

## Modules
This commands help to create a module file support. It takes care of whether the project is inside or outside GOPATH.

```bash
$ gpm mod

# Creating modules file...
gpm: Enter module name: 
github.com/yashishdua/gpm
go: creating new go.mod: module github.com/yashishdua/gpm
```
