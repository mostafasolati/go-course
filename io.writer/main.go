package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type String struct {
	str string
}

func (s *String) Write(p []byte) (int, error) {
	s.str += string(p)
	return len(p), nil
}

func SayHello(w io.Writer) {
	w.Write([]byte("Hello World!"))
}

func main() {
	// var mystr String
	// file, _ := os.Create("hello.txt")
	w := strings.Builder{}
	w.WriteString("This is a beginging")

	SayHello(&w)
	// str = "hello" + "world" + "..."
	// fmt.Println(mystr)
	fmt.Println(w.String())

	b := new(bytes.Buffer)
	b.WriteString("Hello again")
	b.Write([]byte("this is a text"))
	fmt.Fprint(b, "\nHello this is a sample of using Fprint")

	fmt.Fprint(os.Stdout, "This is a text to be shown in terminal", 572, w)

}
