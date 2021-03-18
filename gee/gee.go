package gee

import (
	"log"
	"net/http"
)

// 定义gee使用的请求头
type HandlerFunc func(*Context)

// 包含路由表，路由表中包含操作
type Engine struct {
	router *router
}

// 构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由到路由表中
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

// 根据方法和路由参数写入路由表
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 根据方法和路由参数写入路由表
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 开始运行，监听端口
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 获取请求并处理
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
