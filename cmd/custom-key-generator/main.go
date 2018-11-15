// +build ignore

//this file generates the handlers/custom_keys.go file
//use go generate to update it

package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/bgadrian/fastfaker"
)

func main() {

	if len(os.Args) != 2 {
		log.Panicf("requires one parameter, of ['go', 'markdown']")
	}

	type funct struct {
		Key, Call, Example string
	}
	list := make([]funct, 0, 100)

	faker := fastfaker.NewFastFaker()
	faker.Seed(42)
	fakerType := reflect.TypeOf(faker)
	fakerValue := reflect.ValueOf(faker)

	countOfMethods := fakerType.NumMethod()
	for i := 0; i < countOfMethods; i++ {
		method := fakerType.Method(i)

		//here we only collect simple functions, no parameters
		//the first parameter is the actual struct (because is a method)
		if method.Type.NumIn() > 1 {
			continue
		}

		var example string

		call := "f." + method.Name
		g := byte('g')

		res := fakerValue.Method(i).Call([]reflect.Value{})
		if len(res) != 1 {
			log.Panicf("weird method: %s", fakerValue.Method(i).String())
		}
		switch res[0].Interface().(type) {
		case string:
			example = res[0].String()
			call = fmt.Sprintf("func(f *fastfaker.Faker) string { return %s() }", call)
		case int, int8, int16, int32, int64:
			example = strconv.Itoa(int(res[0].Int()))
			call = fmt.Sprintf("func(f *fastfaker.Faker) string { return strconv.Itoa(int(%s())) }", call)
		case uint8, uint16, uint32, uint64:
			example = strconv.Itoa(int(res[0].Uint()))
			call = fmt.Sprintf("func(f *fastfaker.Faker) string { return strconv.Itoa(int(%s())) }", call)
		case float32:
			example = strconv.FormatFloat(float64(res[0].Float()), 'g', -1, 32)
			call = fmt.Sprintf("func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(%s()), %d, -1, 32) }", call, g)
		case float64:
			example = strconv.FormatFloat(float64(res[0].Float()), 'g', -1, 64)
			call = fmt.Sprintf("func(f *fastfaker.Faker) string { return strconv.FormatFloat(float64(%s()), %d, -1, 64) }", call, g)
		default:
			continue
			//example = fmt.Sprintf("%v", res[0])
			//bytes, err := json.Marshal(res[0])
			//if err != nil {
			//	log.Panic(err)
			//}
			//example = string(bytes)
		}

		list = append(list, funct{
			Key:     strings.ToLower(method.Name),
			Call:    call,
			Example: example,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Key < list[j].Key
	})

	//TODO learn how to use the go generate

	if os.Args[1] == "go" {
		packageTemplate.Execute(os.Stdout, list)
	} else {
		readmeTemplate.Execute(os.Stdout, list)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package handlers

import (
	"github.com/bgadrian/fastfaker"
	"strconv"
)

type fakerer func(*fastfaker.Faker) string
var keys map[string]fakerer

func init() {
	keys = make(map[string]fakerer)
	
	{{range .}}keys["{{.Key}}"] = {{.Call}}
	{{end}}
}
`))

var readmeTemplate = template.Must(template.New("").Parse(`### Custom data 
##### template variables and examples

All the variables that can be used with the **/custom/** entrypoint.
For more info see [Readme](./README.md).

| Variable | Example |
|---------|--------|
{{range .}}| **&#126;{{.Key}}&#126;** | **{{.Example}}** |
{{end}}


> This file is generated automatically at build.

`))
