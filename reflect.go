package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("variable type:", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("variable value:", reflectValue.Int())
	// }

	person := Person{"Raditeo", 23}

	var sampleType = reflect.TypeOf(person)
	var valueType = reflect.ValueOf(person)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType)

	fmt.Println(valueType.Field(0).Interface())
}
