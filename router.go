// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz-mylist/biz/handler"
	"hertz-mylist/biz/router/middleware"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/login", middleware.HzJwtMw.LoginHandler)
}
