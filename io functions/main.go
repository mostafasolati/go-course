package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	w := strings.Builder{}
	w.Write([]byte("Hello World "))
	io.WriteString(&w, "Hello World")

	r := strings.NewReader("this is a sample text sdflksdjflj")
	bs, err := io.ReadAll(r)
	fmt.Println(string(bs), err, len(bs))
	r.Seek(0, io.SeekStart)

	io.CopyN(&w, r, 5)
	fmt.Println(w.String())

	r.Seek(0, io.SeekStart)
	newReader := io.TeeReader(r, os.Stdout)
	io.ReadAll(newReader)

	rr, ww := io.Pipe()
	go func() {
		io.WriteString(ww, "\nWriting to pipe")
		ww.Close()
	}()
	io.Copy(os.Stdout, rr)
	
}
