package controllers

import (
	gs "github.com/0x19/gotostruct"
	"github.com/revel/revel"
	"io/ioutil"
	"strings"
)

func (c App) Parse() revel.Result {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil || len(body) < 2 {
		c.Response.Status = 400
		return c.RenderText(fmt.Sprintf("Could not read json: %v \n", err))
	}

	structName := c.Params.Get("struct")

	if structName == "" {
		structName = "ExampleStruct"
	}

	gts := gs.GotoStruct{Name: structName}

	if response, err := gts.Generate(strings.NewReader(string(body))); err != nil {
		c.Response.Status = 400
		return c.RenderText(fmt.Sprintf("Could not generate struct: %v \n", err))
	} else {
		return c.RenderText(string(response))
	}
}
