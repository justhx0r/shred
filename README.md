[![GoDoc](https://godoc.org/github.com/justhx0r/shred?status.svg)](https://godoc.org/github.com/lu4p/shred)
[![License](https://img.shields.io/github/license/justhx0r/shred.svg)](https://unlicense.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/justhx0r/shred)](https://goreportcard.com/report/github.com/lu4p/shred)
# shred
 Package shred is a golang library to mimic the functionality of the linux `shred` command
 
 ## Usage
```golang
package main
import (
  "github.com/justhx0r/shred"
)

func main(){
	shredconf := shred.Conf{Times: 1, Zeros: true, Remove: false}
	shredconf.Path("filename")
}
```

## Installation
```
go get -u github.com/justhx0r/shred 
```
