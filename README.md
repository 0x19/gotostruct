[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/0x19/gotostruct/tree/master/LICENSE)
[![Build Status](https://travis-ci.org/0x19/gotostruct.svg)](https://travis-ci.org/0x19/gotostruct)
[![Go 1.3 Ready](https://img.shields.io/badge/Go%201.3-Ready-green.svg?style=flat)]()
[![Go 1.4 Ready](https://img.shields.io/badge/Go%201.4-Ready-green.svg?style=flat)]()

GotoStruct
====

Open source [Go](http://golang.org) package designed to help you converting JSON objects into Go Structs.


### Examples

List of examples can be found within [GotoStruct Examples](https://github.com/0x19/gotostruct/tree/master/examples).

Bellow you can see the most common one.

```go
package main

import (
	"fmt"
	gs "github.com/0x19/gotostruct"
	"strings"
)

func main() {
	gts := gs.GotoStruct{
		Name: "ExampleStruct",
	}

	reader := strings.NewReader(`{"simple" : "json", "test": null, "something": [{"baz": 1}, {"baz": 2}]}`)

	response, err := gts.Generate(reader)

	if err != nil {
		fmt.Errorf("Got Error while generating struct: ", err)
	}

	fmt.Println(string(response))
}
```

### Related Work

Parts of [GotoStruct](https://github.com/0x19/gotostruct) code are taken from [Json-to-struct](https://github.com/tmc/json-to-struct).


### Contributing
I encourage you to contribute to GotoStruct! Please check out the [Contributing to GotoStruct](https://github.com/0x19/gotostruct/tree/master/CONTRIBUTING.md) for guidelines about how
to proceed.