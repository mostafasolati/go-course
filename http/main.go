/*
mysite.com/home.html

GET /home.html HTTP/1.1
Host: mysite.com
Cookies:...
Authorization:...

{"first_name":"ali","username:ali.mohammadi"}



HTTP/1.1 200 OK
Date: Sun, 28 Jul 2013 15:37:37 GMT
Server: Apache
Last-Modified: Sun, 07 Jul 2013 06:13:43 GMT Transfer-Encoding: chunked
Connection: Keep-Alive
Content-Type: text/html; charset=UTF-8

Webpage Content
*/

package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func main() {
	// conn, err := net.Dial("tcp", "example.com:80")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Fprint(conn, "GET / HTTP/1.1\r\n")
	// fmt.Fprint(conn, "Host: example.com\r\n")
	// fmt.Fprint(conn, "Connection: close\r\n")
	// fmt.Fprint(conn, "\r\n")

	// response, err := io.ReadAll(conn)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(response))

	listner, err := net.Listen("tcp", "localhost:5050")
	defer listner.Close()
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listner.Accept()
		defer conn.Close()
		if err != nil {
			panic(err)
		}

		t := time.Now()
		buf := new(bytes.Buffer)
		buf.WriteString("<html><body>")
		buf.WriteString(fmt.Sprintf("<h1>%s: %d:%d:%d</h1>", "ساعت", t.Hour(), t.Minute(), t.Second()))
		buf.WriteString("</body></html>")

		lines := []string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/html; charset=UTF-8",
			"Connection: close",
			fmt.Sprintf("Content-Length: %d", buf.Len()), // content-length
			"",
			buf.String(), // body
		}

		for _, line := range lines {
			fmt.Fprint(conn, line+"\r\n")
		}

	}

}
