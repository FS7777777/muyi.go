package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main(){
	fmt.Println(os.Args)
	r := gin.Default()
	r.LoadHTMLGlob("dist/index.html")    // 添加入口index.html
	//r.LoadHTMLFiles("static/*/*")	// 添加资源路径
	r.Static("/assets", "./dist/assets")
	r.Static("/css", "./dist/css")
	r.Static("/img", "./dist/img")
	r.Static("/js", "./dist/js")
	r.Static("/loading", "./dist/loading")
	r.Static("avatar2.jpg", "./dist/avatar2.jpg")
	r.Static("logo.png", "./dist/logo.png")
	r.Static("color.less", "./dist/color.less")
	r.StaticFile("/", "./dist/index.html")

	r.Run(":8080")
}