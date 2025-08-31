// 代码生成时间: 2025-08-31 19:33:08
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/labstack/echo"
    "github.com/mattes/migrate"
    \_ "github.com/mattes/migrate/file"
    \_ "github.com/mattes/migrate/database/postgres"
)

// MigrationFileNamePattern defines the pattern for migration files.
const MigrationFileNamePattern = `^[0-9]+_.*\.sql$`

// MigrationDir is the directory where migration files are stored.
var MigrationDir string

func main() {
    // Initialize the Echo server.
    e := echo.New()

    // Set the directory for migration files.
    // This should be set to the actual directory in your environment.
    MigrationDir = os.Getenv("MIGRATION_DIR")
    if MigrationDir == "" {
        log.Fatal("MIGRATION_DIR environment variable is not set.")
    }

    // Define the route for triggering migrations.
    e.GET("/migrate", migrateHandler)

    // Start the server.
    e.Logger.Fatal(e.Start(":8080"))
}

// migrateHandler is the handler function for the migration route.
func migrateHandler(c echo.Context) error {
    // Create a new migration instance.
    m, err := migrate.NewFileSystem(
        filepath.Join(MigrationDir, ":%v"),
        "github.com/mattes/migrate/database/postgres",
    )
    if err != nil {
        return err
    }

    // Perform the migration.
    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("migration failed: %w", err)
    }

    // Return a success message.
    return c.JSON(http.StatusOK, map[string]string{
        "status": "migration successful",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}
