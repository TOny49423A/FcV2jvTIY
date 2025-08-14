// 代码生成时间: 2025-08-14 11:56:56
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "encoding/json"
    "log"
)

// MathOperation 结构体定义了数学操作的输入参数
type MathOperation struct {
    A float64 `json:"a"`
    B float64 `json:"b"`
}

// Result 结构体定义了数学操作的结果
type Result struct {
    Result float64 `json:"result"`
    Error  string  `json:"error"`
}

// MathHandler 定义了一个处理数学运算的处理器
func MathHandler(c echo.Context) error {
    var op MathOperation
    if err := json.NewDecoder(c.Request().Body).Decode(&op); err != nil {
        return c.JSON(http.StatusBadRequest, Result{Result: 0, Error: err.Error()})
    }

    switch c.Request().Method {
    case http.MethodPost:
        // 处理加法运算
        if c.Path() == "/add" {
            return c.JSON(http.StatusOK, Result{Result: op.A + op.B, Error: ""})
        }
        // 处理减法运算
        if c.Path() == "/subtract" {
            return c.JSON(http.StatusOK, Result{Result: op.A - op.B, Error: ""})
        }
        // 处理乘法运算
        if c.Path() ==="/multiply" {
            return c.JSON(http.StatusOK, Result{Result: op.A * op.B, Error: ""})
        }
        // 处理除法运算
        if c.Path() == "/divide" {
            if op.B == 0 {
                return c.JSON(http.StatusBadRequest, Result{Result: 0, Error: "division by zero"})
            }
            return c.JSON(http.StatusOK, Result{Result: op.A / op.B, Error: ""})
        }
    }
    
    return c.JSON(http.StatusMethodNotAllowed, Result{Result: 0, Error: "method not allowed"})
}

func main() {
    e := echo.New()
    
    // 添加路由
    e.POST("/add", MathHandler)
    e.POST("/subtract", MathHandler)
    e.POST="/multiply", MathHandler)
    e.POST("/divide", MathHandler)

    // 启动服务器
    log.Fatal(e.Start(":8080"))
}
