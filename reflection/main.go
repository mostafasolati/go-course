package main

import (
	"errors"
	"fmt"
	"reflect"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func json(input any) (string, error) {
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Struct {
		return "", errors.New("unsupported " + t.Kind().String() + " type")
	}

	count := v.NumField()
	ret := ""
	for i := range count {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		switch fieldType.Type.Kind() {
		case reflect.Int:
			ret += fmt.Sprintf(`"%s":%d`, fieldType.Name, fieldValue.Int())
		case reflect.String:
			ret += fmt.Sprintf(`"%s":"%s"`, fieldType.Name, fieldValue.String())
		}

		if i < count-1 {
			ret += ","
		}
	}
	return "{" + ret + "}", nil
}

/*
{
	"FirstName" : "Ebrahim",
	"LastName" : "Hamedi",
	"Age" : 80
}

*/

func main() {
	u := User{FirstName: "Ebrahim", LastName: "Hamedi", Age: 80}
	// u := []int{1, 4, 7}
	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)
	fmt.Println("Value of u:", v)
	fmt.Println("Type of u:", t)
	fmt.Println("Kind of u:", t.Kind())

	if t.Kind() == reflect.Slice {
		numbers := v.Interface().([]int)
		for _, num := range numbers {
			fmt.Println(num)
		}
	}

	n := 5
	fmt.Println("n Before reflection:", n)
	v = reflect.ValueOf(&n).Elem()
	v.SetInt(100)
	fmt.Println("n After reflection:", n)

	ret, err := json(u)
	fmt.Println(ret, err)
}
