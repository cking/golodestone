golodestone
====
[![GoDoc](https://godoc.org/github.com/cking/golodestone?status.svg)](https://godoc.org/github.com/cking/golodestone)
[![GoWalker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com/cking/golodestone)
[![Go report](http://goreportcard.com/badge/cking/golodestone)](http://goreportcard.com/report/cking/golodestone)
[![MS-PL License](https://img.shields.io/github/license/cking/golodestone.svg)](https://github/cking/golodestone/blob/master/LICENSE)
[![Build status](https://img.shields.io/travis/cking/golodestone.svg)](https://travis-ci.org/cking/golodestone)

golodestone is an API to interface with Final Fantasy XIV Lodestone. It fetches the required data by scraping the official Lodestone pages.

## Getting Started

The basic interface is all stored in `golodestone`, to interface with the Characters or Free Companies, import the relevant sub module (`golodestone/character` and `golodestone/freecompany`)

### Usage

Import the package into your project.

```go
import "github.com/cking/golodestone"
import "github.com/cking/golodestone/character"
import "github.com/cking/golodestone/freecompany"
```

Query the data you want to see.
```go
characters, err := character.SearchCharacter("Character Name")
```
