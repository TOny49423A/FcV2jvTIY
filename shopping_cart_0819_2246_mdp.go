// 代码生成时间: 2025-08-19 22:46:35
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

// ShoppingCart 购物车结构体
type ShoppingCart struct {
    Items map[string]int
}

// CartService 购物车服务
type CartService struct {
    carts map[string]ShoppingCart
}

// NewCartService 创建新的购物车服务实例
func NewCartService() *CartService {
    return &CartService{
        carts: make(map[string]ShoppingCart),
    }
}

// AddItem 向购物车中添加商品
func (s *CartService) AddItem(userID string, itemID string, quantity int) error {
    if _, exists := s.carts[userID]; !exists {
        s.carts[userID] = ShoppingCart{Items: make(map[string]int)}
    }
    s.carts[userID].Items[itemID] += quantity
    return nil
}

// RemoveItem 从购物车中移除商品
func (s *CartService) RemoveItem(userID string, itemID string) error {
    if cart, exists := s.carts[userID]; exists {
        if _, exists := cart.Items[itemID]; exists {
            delete(cart.Items, itemID)
            return nil
        }
    }
    return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
}

// GetCart 获取用户的购物车
func (s *CartService) GetCart(userID string) (ShoppingCart, error) {
    if cart, exists := s.carts[userID]; exists {
        return cart, nil
    }
    return ShoppingCart{}, echo.NewHTTPError(http.StatusNotFound, "Cart not found")
}

// SetupRoutes 设置路由
func SetupRoutes(e *echo.Echo, service *CartService) {
    e.GET("/cart/:userID", func(c echo.Context) error {
        userID := c.Param("userID")
        cart, err := service.GetCart(userID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, cart)
    })

    e.POST("/cart/:userID/add", func(c echo.Context) error {
        userID := c.Param("userID")
        var payload struct {
            ItemID string `json:"itemID"`
            Quantity int `json:"quantity"`
        }
        if err := c.Bind(&payload); err != nil {
            return err
        }
        if err := service.AddItem(userID, payload.ItemID, payload.Quantity); err != nil {
            return err
        }
        return c.NoContent(http.StatusOK)
    })

    e.POST("/cart/:userID/remove", func(c echo.Context) error {
        userID := c.Param("userID")
        var payload struct {
            ItemID string `json:"itemID"`
        }
        if err := c.Bind(&payload); err != nil {
            return err
        }
        if err := service.RemoveItem(userID, payload.ItemID); err != nil {
            return err
        }
        return c.NoContent(http.StatusOK)
    })
}

func main() {
    e := echo.New()
    service := NewCartService()
    SetupRoutes(e, service)
    e.Logger.Fatal(e.Start(":8080"))
}