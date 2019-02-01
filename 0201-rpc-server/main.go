package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func hello2(name string) string {
	return "Hello222 " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	service.AddFunction("hello2", hello2)

	router := gin.Default()
	router.Any("/path", func(c *gin.Context) {
		service.ServeHTTP(c.Writer, c.Request)
	})
	router.Run(":8080")
}
