// 代码生成时间: 2025-08-05 14:57:29
package main

import (
    "fmt"
    "sort"
    "math/rand"
    "time"
    "net/http"

    "github.com/labstack/echo"
)

// SortAlgorithmHandler defines the handler for sorting algorithms.
func SortAlgorithmHandler(c echo.Context) error {
    // Generate a random slice of integers.
    numbers := generateRandomSlice(100)
    fmt.Println("Original slice: ", numbers)

    // Sort the slice using sort.Ints.
    fmt.Println("Sorted slice using sort.Ints: ", sort.Ints(numbers))

    // Sort the slice in descending order.
    fmt.Println("Sorted slice in descending order: ", sort.Reverse(sort.IntSlice(numbers)))

    // Return a success message as JSON.
    return c.JSON(http.StatusOK, map[string]string{
        "message": "Sorting algorithms executed successfully.",
    })
}

// generateRandomSlice generates a random slice of integers of a given size.
func generateRandomSlice(size int) []int {
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int, size)
    for i := range numbers {
        numbers[i] = rand.Intn(1000)
    }
    return numbers
}

func main() {
    e := echo.New()
    e.GET("/sort", SortAlgorithmHandler)
    e.Logger.Fatal(e.Start(":8080"))
}
