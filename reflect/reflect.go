package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(1)
	fmt.Println("x :", v.Elem().Interface())
}