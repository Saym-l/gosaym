package main

import (
	"net/http"
	"saym"
)

func main() {
	r := saym.New()
	r.GET("/", func(ctx *saym.Context) {
		ctx.Html(http.StatusOK, "<h1>hello saym</h1>")
	})
	r.GET("/hello", func(ctx *saym.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.POST("/login", func(ctx *saym.Context) {
		ctx.JSON(http.StatusOK, saym.H{
			"name": ctx.PostForm("name"),
			"pwd":  ctx.PostForm("pwd"),
		})
	})
	r.Run(":8888")
}
