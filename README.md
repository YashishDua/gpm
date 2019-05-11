# gpm (Go Package Manager)

![Image](gpm.png)

gpm is a tool for managing Go projects.

# Overview

gpm provides commands:
* [To create boilerplate project directory](#create)
* [To build project either using vendor or modules (inside/outside GOPATH)](#build)
* [To update Go version](#version)
* [To add vendor support using modules](#vendor)
* [To add module support (inside/outside GOPATH)](#modules)

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

## Update
This commands updates Go version to specified version. If version not specified, uses default 1.12.5.

```bash
$ gpm update -v=1.12.1

# Updating Go version...
gpm: Uninstalling previous version
gpm: Download go1.12.5.darwin-amd64.tar.gz binary
gpm: Extracting Go archive
gpm: Go updated successfuly
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
