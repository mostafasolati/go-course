package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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
	fmt.Println("String.Read called:", n, "bytes read")
	return n, err
}

func (s *String) Write(b []byte) (int, error) {
	s.data += string(b)
	fmt.Println("String.Write called:", len(b), "bytes wrote")
	return len(b), nil
}

func main() {
	fmt.Println("Without bufio:")
	rw := &String{}
	fmt.Fprint(rw, "This")
	fmt.Fprint(rw, "is")
	fmt.Fprint(rw, "a")
	fmt.Fprint(rw, "sample")
	fmt.Fprint(rw, "text")

	buf := make([]byte, 2)
	var err error

	for err != io.EOF {
		_, err = rw.Read(buf)
	}

	fmt.Println("With bufio:")
	rw = &String{}
	w := bufio.NewWriter(rw)
	fmt.Fprint(w, "This")
	fmt.Fprint(w, "is")
	fmt.Fprint(w, "a")
	fmt.Fprint(w, "sample")
	fmt.Fprint(w, "text")
	w.Flush()

	r := bufio.NewReader(rw)
	buf = make([]byte, 4)
	err = nil
	var n int

	for err != io.EOF {
		n, err = r.Read(buf)
		fmt.Println(string(buf[:n]))
	}

	sr := strings.NewReader("This is, line 1\nThi,s is line2\nThis is ,line3")
	scanner := bufio.NewScanner(sr)
	scanner.Split(splitByComma)
	for scanner.Scan() {
		fmt.Println("Text:", scanner.Text())
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("Your input is:", input.Text())
}

func splitByComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
