// 代码生成时间: 2025-08-05 08:35:59
package main

import (
    "net/http"
    "strings"
    "gopkg.in/go-playground/validator.v9"
    "github.com/labstack/echo"
)

// Form represents the data that will be validated
type Form struct {
    Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"gte=1,lte=130"`
}

// ValidateForm validates the form data
func ValidateForm(c echo.Context) error {
    f := new(Form)
    if err := c.Bind(f); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }
    
    validate := validator.New()
    if err := validate.Struct(f); err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            return echo.NewHTTPError(http.StatusBadRequest, err.Namespace() + " is invalid: " + err.ActualTag())
        }
    }
    
    // If no errors, continue processing
    return c.JSON(http.StatusOK, f)
}

func main() {
    e := echo.New()
    
    // Define route for form submission
    e.POST("/form", ValidateForm)
    
    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}
