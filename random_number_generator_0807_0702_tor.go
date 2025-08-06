// 代码生成时间: 2025-08-07 07:02:21
package main

import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
    "log"
    "math/big"

    "github.com/labstack/echo"
)

// RandomNumberGeneratorHandler returns a random number between 1 and 100
func RandomNumberGeneratorHandler(c echo.Context) error {
    randomNumber, err := GenerateRandomNumber()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error generating random number: %v", err))
    }
    return c.JSON(http.StatusOK, map[string]int{
        "randomNumber": randomNumber,
    })
}

// GenerateRandomNumber generates a random number between 1 and 100
func GenerateRandomNumber() (int64, error) {
    max := big.NewInt(100)
    randomNumber, err := rand.Int(rand.Reader, max)
    if err != nil {
        return 0, fmt.Errorf("failed to generate random number: %w", err)
    }
    return randomNumber.Int64() + 1, nil // Add 1 to include 1 in the range
}

func main() {
    e := echo.New()
    e.GET("/random", RandomNumberGeneratorHandler)

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
