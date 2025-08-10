// 代码生成时间: 2025-08-11 01:47:52
// database_migration_tool.go

package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/go-pg/migrations/v8"
    "github.com/labstack/echo/v4"
    _ "github.com/go-pg/pg/v10/orm" // PostgreSQL driver
)

// DatabaseConfig contains database connection information
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Database string
}

func main() {
    // Define Echo instance
    e := echo.New()

    // Define database configuration
    dbConfig := DatabaseConfig{
        Host:     "localhost",
        Port:     5432,
        User:     "user",
        Password: "password",
        Database: "database",
    }

    // Initialize database connection
    db := initDB(dbConfig)
    defer db.Close()

    // Register migration handler
    e.GET("/migrate", func(c echo.Context) error {
        if err := migrateDatabase(db); err != nil {
            return err
        }
        return c.JSON(echo.StatusOK, map[string]string{
            "message": "Database migration successful",
        })
    })

    // Start the Echo server
    e.Logger.Fatal(e.Start(":8080"))
}

// initDB initializes the database connection
func initDB(config DatabaseConfig) *sql.DB {
    pgOptions := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
        config.User, config.Password, config.Host, config.Port, config.Database)
    db, err := sql.Open("pg", pgOptions)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    // Ensure the database connection is live
    if err = db.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }
    return db
}

// migrateDatabase applies database migrations
func migrateDatabase(db *sql.DB) error {
    // Define migrations
    migrationsPath := filepath.Join("migrations", "*.sql\)
    runner, err := migrations.NewMigrationRunner(migrationsPath, db)
    if err != nil {
        return fmt.Errorf("failed to create migration runner: %w", err)
    }
    // Run migrations
    if err := runner.Up(); err != nil {
        return fmt.Errorf("migration failed: %w", err)
    }
    return nil
}
