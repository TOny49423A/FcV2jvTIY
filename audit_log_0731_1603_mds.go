// 代码生成时间: 2025-07-31 16:03:55
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "github.com/labstack/echo"
)

// AuditLog represents a security audit log record
type AuditLog struct {
    Timestamp time.Time `json:"timestamp"`
    Action    string    `json:"action"`
    UserID    string    `json:"userId"`
    IP        string    `json:"ip"`
    Hash      string    `json:"hash"`
}

// NewAuditLog creates a new audit log record
func NewAuditLog(action string, userID string, ip string) *AuditLog {
    hash := sha1Hash(action, userID, ip)
    return &AuditLog{
        Timestamp: time.Now(),
        Action:    action,
        UserID:    userID,
        IP:        ip,
        Hash:      hash,
    }
}

// sha1Hash creates a SHA1 hash of the action, userID, and IP
func sha1Hash(action, userID, ip string) string {
    sha1 := sha1.New()
    sha1.Write([]byte(action + userID + ip))
    return hex.EncodeToString(sha1.Sum(nil))
}

// LogAudit writes the audit log record to a file
func LogAudit(log *AuditLog) error {
    f, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    
    encoder := json.NewEncoder(f)
    err = encoder.Encode(log)
    if err != nil {
        return err
    }
    return nil
}

// AuditLogHandler handles the audit log generation
func AuditLogHandler(c echo.Context) error {
    action := c.QueryParam("action")
    userID := c.QueryParam("userId")
    ip := c.Request().RemoteAddr
    
    if action == "" || userID == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing action or userId")
    }
    
    auditLog := NewAuditLog(action, userID, ip)
    if err := LogAudit(auditLog); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to log audit")
    }
    
    return c.JSON(http.StatusOK, auditLog)
}

func main() {
    e := echo.New()
    e.GET("/log", AuditLogHandler)
    
    e.Logger.Fatal(e.Start(":8080"))
}
