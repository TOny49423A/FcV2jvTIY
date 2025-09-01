// 代码生成时间: 2025-09-01 12:01:59
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "log"
)

// SortingAlgorithmHandler is the handler function for sorting algorithm requests.
func SortingAlgorithmHandler(c echo.Context) error {
    // Example integer slice to sort.
    numbers := []int{5, 2, 9, 1, 5, 6}
    
    // Sorting the slice using the built-in sort function.
    sort.Ints(numbers)
    
    // Return the sorted slice as a JSON response.
    return c.JSON(http.StatusOK, numbers)
}

func main() {
    // Create a new Echo instance.
    e := echo.New()
    
    // Define the route for the sorting algorithm.
    e.GET("/sort", SortingAlgorithmHandler)
    
    // Start the Echo server.
    log.Println("Server is running on :8080")
    if err := e.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}