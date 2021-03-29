package main

import (
	"gee"
	"net/http"
)

var test1 = func(c *gee.Context) {
	// /hello?name=geektutu
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

var test2 = func(c *gee.Context) {
	// /hello/geektutu
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
}

var test3 = func(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
}