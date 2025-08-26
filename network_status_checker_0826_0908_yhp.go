// 代码生成时间: 2025-08-26 09:08:51
package main

import (
	"net"
	"time"
	
	"github.com/labstack/echo"
)

// NetworkConnectionStatus is the struct that holds the status of the network connection.
type NetworkConnectionStatus struct {
	Status string `json:"status"`
}

// checkNetworkConnection attempts to establish a TCP connection to a given host and port.
// It returns a boolean indicating whether the connection was successful or not.
func checkNetworkConnection(host string, port int) bool {
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Second*10)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// NetworkStatusHandler is a handler function that checks network status and returns the result.
func NetworkStatusHandler(c echo.Context) error {
	host := c.QueryParam("host")
	port := c.QueryParam("port")
	
	// Validate the input parameters
	if host == "" || port == "" {
		return c.JSON(echo.StatusBadRequest, NetworkConnectionStatus{Status: "Invalid input parameters"})
	}
	
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return c.JSON(echo.StatusBadRequest, NetworkConnectionStatus{Status: "Invalid port number"})
	}
	
	// Check network connection status
	isConnected := checkNetworkConnection(host, portInt)
	status := "Connected"
	if !isConnected {
		status = "Disconnected"
	}
	
	return c.JSON(echo.StatusOK, NetworkConnectionStatus{Status: status})
}

func main() {
	e := echo.New()
	e.GET("/network-status", NetworkStatusHandler)
	e.Logger.Fatal(e.Start(":8080"))
}