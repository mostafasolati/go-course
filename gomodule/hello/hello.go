package hello

import "fmt"

func SayHello(name string) {
	fmt.Println(greeting(name))
}

func greeting(name string) string {
	return "Hello " + name
}

func init() {
	fmt.Println("Init strated")
}

func init() {
	fmt.Println("Init2 started")
}
