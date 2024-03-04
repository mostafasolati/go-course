package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed file1.txt
var file1 embed.FS

//go:embed images
var folder embed.FS

func main() {
	// f, err := os.Open("file1.txt")
	f, err := file1.Open("file1.txt")
	// defer f.Close()
	if err != nil {
		panic(err)
	}
	bs, _ := io.ReadAll(f)
	fmt.Println(string(bs))

	// fmt.Println(file1)

	files, _ := folder.ReadDir("images")
	for _, file := range files {
		fmt.Println(file.Name())
		f, _ = folder.Open("images/" + file.Name())
		bs, _ := io.ReadAll(f)
		fmt.Println(string(bs))
	}

}
