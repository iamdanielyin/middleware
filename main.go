package main

import "fmt"

type HandlerFunc func(*Context)

type Context struct {
	index    int
	handlers []HandlerFunc
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func main() {
	c := &Context{
		index: -1,
		handlers: []HandlerFunc{
			func(c *Context) {
				fmt.Println("middleware1-start")
				c.Next()
				fmt.Println("middleware1-end")
			},
			func(c *Context) {
				fmt.Println("-----middleware2-start")
				c.Next()
				fmt.Println("-----middleware2-end")
			},
			func(c *Context) {
				fmt.Println("---------middleware3-start")
				c.Next()
				fmt.Println("---------middleware3-end")
			},
			func(c *Context) {
				fmt.Println("------------------" + "hello world")
			},
		},
	}
	c.Next()

}
