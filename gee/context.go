package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)
/*
   抽象出上下文，包含writer，req，path，method，statuscode等信息
   作为处理的整个流程的处理件
*/

type H map[string]interface{}

// 上下文结构体
type Context struct {
	// 来源数据
	Writer http.ResponseWriter
	Req *http.Request

	// 请求信息
	Path string
	Method string

	// 回复信息
	StatusCode int
}

// 构造上下文
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context {
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

// 根据键值获取postform
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 获取查询
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置response状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 回复字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 回复JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer) // 通过c.Writer构造json encode
	if err := encoder.Encode(obj); err != nil { // 将obj接口转换json后写入c.Writer
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 回复数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 回复HTML
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
