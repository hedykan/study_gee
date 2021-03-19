package main

import (
	"gee"
	"net/http"
)

func main() {
	result := gee.New()

	// 把方法写入路由
	test1 := func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	}
	result.GET("/", test1)
	test2 := func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	}
	result.GET("/hello", test2)
	test3 := func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	}
	result.POST("/login", test3)

	result.Run(":9999")
}
