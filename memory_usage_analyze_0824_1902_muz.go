// 代码生成时间: 2025-08-24 19:02:02
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/exec"

    "github.com/labstack/echo/v4" // Import the Echo framework
    "github.com/labstack/echo/v4/middleware"
)

// MemoryUsage represents the memory usage data structure
type MemoryUsage struct {
    Total     int64  `json:"total"`    // Total memory in bytes
    Available int64  `json:"available"` // Available memory in bytes
    Used      int64  `json:"used"`     // Used memory in bytes
    Free      int64  `json:"free"`      // Free memory in bytes
    Active    int64  `json:"active"`   // Active memory in bytes
    Passive   int64  `json:"passive"`  // Passive memory in bytes
    // Add more fields if needed
}

// getMemoryUsage retrieves the memory usage statistics using the 'free' command
func getMemoryUsage() (*MemoryUsage, error) {
    cmd := exec.Command("sh", "-c", "free -m")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return nil, err
    }

    // Parse the output to extract memory usage
    var mu MemoryUsage
    lines := strings.Split(strings.TrimSpace(string(output)), "
")
    if len(lines) < 2 {
        return nil, fmt.Errorf("unexpected output format")
    }

    // Skip the header line
    fields := strings.Fields(lines[1])
    if len(fields) < 6 {
        return nil, fmt.Errorf("unexpected output format")
    }

    mu.Total, _ = strconv.Atoi(fields[1])
    mu.Used, _ = strconv.Atoi(fields[2])
    mu.Free, _ = strconv.Atoi(fields[3])
    mu.Shared, _ = strconv.Atoi(fields[4])
    mu.Buffer, _ = strconv.Atoi(fields[5])
    mu.Available = mu.Free + mu.Buffer + mu.Shared
    mu.Active = mu.Used
    mu.Passive = mu.Available - mu.Free
    return &mu, nil
}

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/memory", func(c echo.Context) error {
        mu, err := getMemoryUsage()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(http.StatusOK, mu)
    })

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}