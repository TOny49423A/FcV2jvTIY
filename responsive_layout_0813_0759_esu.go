// 代码生成时间: 2025-08-13 07:59:06
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "html/template"
    "log"
)

// LayoutTemplate 存储布局的HTML模板
var LayoutTemplate *template.Template

func init() {
    // 编译HTML模板
    var err error
    LayoutTemplate, err = template.ParseGlob("templates/*.html")
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }
}

// Layout 结构体，用于传递数据到模板
type Layout struct {
    Title string
    Body  template.HTML
}

func main() {
    e := echo.New()
    e.Static("/static", "static") // 静态文件服务

    // 路由设置
    e.GET("/", indexHandler)
    e.GET("/about", aboutHandler)
    
    // 启动Echo服务器
    e.Logger.Fatal(e.Start(":8080"))
}

// indexHandler 首页处理器
func indexHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", Layout{
        Title: "Home",
        Body:  "<p>This is the home page.</p>",
    })
}

// aboutHandler 关于页面处理器
func aboutHandler(c echo.Context) error {
    return c.Render(http.StatusOK, "about.html", Layout{
        Title: "About",
        Body:  "<p>This is the about page.</p>",
    })
}