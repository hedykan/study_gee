package main

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/hello", test1)
	r.GET("/hello/:name", test2)
	r.GET("/assets/*filepath", test3)

	r.Run(":9999")
}
