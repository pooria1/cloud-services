package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

func newfunc() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	http.Handle("/", m)
}

func main() {
	fmt.Println("Hello world!")
	go newfunc()
	time.Sleep(time.Second * 20)
}
