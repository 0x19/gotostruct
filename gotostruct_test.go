package gotostruct

import (
	"strings"
	"testing"
)

const (
	SIMPLE_JSON  = `{"simple" : "json", "test": null, "something": [{"baz": 1}, {"baz": 2}]}`
	COMPLEX_JSON = `{
	    "_id": "54b785aa6eeb4f15d81e8d24",
	    "index": 0,
	    "guid": "4ee137bb-16d2-42ab-a979-e0a390350454",
	    "isActive": false,
	    "balance": "$3,366.77",
	    "picture": "http://placehold.it/32x32",
	    "age": 35,
	    "eyeColor": "blue",
	    "name": "Susana Cline",
	    "gender": "female",
	    "company": "ACIUM",
	    "email": "susanacline@acium.com",
	    "phone": "+1 (837) 561-3582",
	    "address": "948 Stuart Street, Kingstowne, Kansas, 1199",
	    "about": "Aute consectetur enim elit amet laborum minim in laboris incididunt. Sint aliquip dolore laboris qui aliqua incididunt aute proident laborum tempor ullamco non. Occaecat proident ad dolore sit laborum. Culpa nisi excepteur nostrud non dolore aute id commodo sit id nostrud. Ipsum aute deserunt nisi velit aliqua pariatur incididunt commodo. Veniam irure occaecat nisi aliqua veniam deserunt irure consectetur excepteur nostrud eu adipisicing non. Cupidatat cupidatat id consequat magna nisi est occaecat.\r\n",
	    "registered": "2014-08-11T09:49:35 -02:00",
	    "latitude": 89.760396,
	    "longitude": 45.684642,
	    "tags": [
	      "fugiat",
	      "esse",
	      "incididunt",
	      "veniam",
	      "aute",
	      "fugiat",
	      "anim"
	    ],
	    "friends": [
	      {
	        "id": 0,
	        "name": "Langley Shannon"
	      },
	      {
	        "id": 1,
	        "name": "Perkins Stone"
	      },
	      {
	        "id": 2,
	        "name": "Elizabeth Flynn"
	      }
	    ],
	    "greeting": "Hello, Susana Cline! You have 9 unread messages.",
	    "favoriteFruit": "strawberry",
	    "test.o-q": null
	}`
)

func TestStructName(t *testing.T) {
	gos := GotoStruct{}

	if gos.Name != "" {
		t.Fatal("Name should be empty and its not")
	}

	gos.SetName("TestNewStructName")

	if gos.Name != "TestNewStructName" {
		t.Fatal("Name should equal TestNewStructName and it's not. - ", gos.Name)
	}
}

func TestSimpleJson(t *testing.T) {
	gos := GotoStruct{Name: "TestStruct"}

	reader := strings.NewReader(SIMPLE_JSON)

	response, err := gos.Generate(reader)

	if err != nil {
		t.Fatal("Got Error while generating struct: ", err)
	}

	if !strings.HasPrefix(string(response), "type TestStruct struct {") {
		t.Fatal("Struct is not starting as it should: ", string(response))
	}
}
