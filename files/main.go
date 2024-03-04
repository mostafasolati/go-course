package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("myfile.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	file := bufio.NewWriter(f)
	file.Write([]byte("line1\n"))
	file.WriteString("text string\n")
	io.WriteString(file, "third line\n")
	fmt.Fprint(file, "text", 5+6, "\n")
	file.Flush()

	err = os.WriteFile("myfile2.txt", []byte("This is a sample text"), 0777)
	if err != nil {
		panic(err)
	}

	f, err = os.Open("myfile2.txt")
	bs, err := io.ReadAll(f)
	fmt.Println("myfile2.txt:", string(bs))

	bs, err = os.ReadFile("myfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("myfile.tx:", string(bs))

	f, err = os.OpenFile("myfile3.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	bs, err = io.ReadAll(f)
	n, err := f.WriteString("Hello World")
	fmt.Println(n, err)
	fmt.Println(string(bs))

	f.Seek(6, io.SeekStart)
	f.Write([]byte("w"))

	f.WriteAt([]byte("oooooooo "), 4)

	os.Remove("myfile4.txt")

	_, err = os.Stat("myfile4.txt")
	if os.IsNotExist(err) {
		fmt.Println("File not found!")
	}

	src, _ := os.Open("myfile.txt")
	dst, _ := os.Create("new_file.txt")
	defer src.Close()
	defer dst.Close()
	io.Copy(dst, src)

	err = os.Mkdir("myfolder2", 0777)
	fmt.Println("Makedir err:", err)
	os.MkdirAll("a/b/c/d", 0777)

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("[folder] ")
		} else {
			fmt.Print("[file] ")
		}
		fmt.Println(file.Name())
	}
}
