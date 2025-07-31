// 代码生成时间: 2025-07-31 21:15:23
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "strings"

    "github.com/labstack/echo/v4" // Importing Echo Framework
)

// HashCalculator defines the structure to hold the Echo instance.
type HashCalculator struct {
    e *echo.Echo
}

// NewHashCalculator creates a new instance of HashCalculator.
func NewHashCalculator() *HashCalculator {
    e := echo.New()
    return &HashCalculator{
        e: e,
    }
}

// CalculateHash handles the HTTP request to calculate the SHA-256 hash of a given input.
func (h *HashCalculator) CalculateHash() echo.HandlerFunc {
    return func(c echo.Context) error {
        input := c.QueryParam("input")
        if input == "" {
            return c.JSON(http.StatusBadRequest, echo.Map{"error": "Input parameter is missing."})
        }

        // Calculate the SHA-256 hash of the input.
        hash := sha256.Sum256([]byte(input))
        hexHash := hex.EncodeToString(hash[:])

        // Return the calculated hash as a JSON response.
        return c.JSON(http.StatusOK, echo.Map{"hash": hexHash})
    }
}

func main() {
    // Create a new instance of HashCalculator.
    hc := NewHashCalculator()

    // Define the route for calculating hash.
    hc.e.GET("/hash", hc.CalculateHash())

    // Start the Echo server.
    hc.e.Start(":8080")
}
