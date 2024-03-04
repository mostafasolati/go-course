package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type ReadWriter interface {
	io.Reader
	io.Writer
}

type String struct {
	data string
	pos  int
}

func (s *String) Read(b []byte) (int, error) {
	n := copy(b, s.data[s.pos:])
	s.pos += n
	var err error
	if s.pos >= len(s.data) {
		err = io.EOF
	}
	return n, err
}

func (s *String) Write(b []byte) (int, error) {
	s.data += string(b)
	return len(b), nil
}

func main() {
	var s String
	s.Write([]byte("Hello World!"))
	var err error

	for err != io.EOF {
		buf := make([]byte, 3)
		_, err = s.Read(buf)
		// fmt.Println(string(buf), n, err)
		fmt.Print(string(buf))
	}
	fmt.Println()

	buf := make([]byte, 20)
	os.Stdin.Read(buf)
	fmt.Println("input:", string(buf))

	sr := strings.NewReader("This is a sample text")
	sr.Read(buf)

	br := bytes.NewReader([]byte("this is a bytes slice"))
	br.Read(buf)

}
