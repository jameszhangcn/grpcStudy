package main

import (
	"errors"
	"fmt"
	"reflect"
)

type MyStruct struct {
	i int
	s string
}

func foo0() int {
	fmt.Println("runing foo0: ")
	return 100
}

func foo1(a int) (string, string) {
	fmt.Println("running foo1: ", a)
	return "aaaa", "bbb"
}

func foo2(a, b int, c string) MyStruct {
	fmt.Println("running foo2:", a, b, c)
	return MyStruct{10, "ccc"}
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {

	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {

	funcs := map[string]interface{}{
		"foo":  foo0,
		"foo1": foo1,
		"foo2": foo2,
	}

	if result, err := Call(funcs, "foo"); err == nil {
		for _, r := range result {
			fmt.Printf(" return: type=%v, value=[%d]\n", r.Type(), r.Int())
		}
	}

	if result, err := Call(funcs, "foo1", 10); err == nil {
		for _, r := range result {
			fmt.Printf(" return: type=%v, value=[%d]\n", r.Type(), r.String())
		}
	}

	if result, err := Call(funcs, "foo2", 20, 30, "test"); err == nil {
		for _, r := range result {
			fmt.Printf(" return: type=%v, value=[%d]\n", r.Type(), r.Interface().(MyStruct))
		}
	}
}
