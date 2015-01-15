package gotostruct

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

var uppercaseFixups = map[string]bool{"id": true, "url": true}

type GotoStruct struct {
	Name string
}

func (g *GotoStruct) SetName(n string) {
	g.Name = n
}

func (g *GotoStruct) Generate(input io.Reader) ([]byte, error) {

	if g.Name == "" {
		g.SetName("ExampleStruct")
	}

	var jsonobj interface{}
	var result map[string]interface{}

	if err := json.NewDecoder(input).Decode(&jsonobj); err != nil {
		return nil, err
	}

	switch jsonobj := jsonobj.(type) {
	case map[string]interface{}:
		result = jsonobj
	case []map[string]interface{}:
		if len(jsonobj) < 1 {
			return nil, fmt.Errorf("empty array")
		}
		result = jsonobj[0]
	default:
		return nil, fmt.Errorf("unexpected type: %T", jsonobj)
	}

	src := fmt.Sprintf("type %s %s}", g.Name, g.build(result, 0))

	formatted, err := format.Source([]byte(src))

	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}
	return formatted, err
}

func (g *GotoStruct) fmtFieldName(s string) string {
	parts := strings.Split(s, "_")

	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	if len(parts) > 0 {
		last := parts[len(parts)-1]
		if uppercaseFixups[strings.ToLower(last)] {
			parts[len(parts)-1] = strings.ToUpper(last)
		}
	}

	assembled := strings.Join(parts, "")
	runes := []rune(assembled)

	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}

	return string(runes)
}

func (g *GotoStruct) typeForValue(value interface{}) string {
	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)

		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}

		if len(types) == 1 {
			return "[]" + g.typeForValue(objects[0])
		}

		return "[]interface{}"
	} else if object, ok := value.(map[string]interface{}); ok {
		return g.build(object, 0) + "}"
	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}

	return reflect.TypeOf(value).Name()
}

func (g *GotoStruct) build(obj map[string]interface{}, depth int) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))

	for key := range obj {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := obj[key]
		valueType := g.typeForValue(value)

		//If a nested value, recurse
		switch value := value.(type) {
		case []map[string]interface{}:
			valueType = "[]" + g.build(value[0], depth+1) + "}"
		case map[string]interface{}:
			valueType = g.build(value, depth+1) + "}"
		}

		fieldName := g.fmtFieldName(key)
		structure += fmt.Sprintf("\n%s %s `json:\"%s\"`", fieldName, valueType, key)
	}

	return structure
}
