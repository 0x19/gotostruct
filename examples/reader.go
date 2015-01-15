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
