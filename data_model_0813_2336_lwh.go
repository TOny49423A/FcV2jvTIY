// 代码生成时间: 2025-08-13 23:36:20
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
    "github.com/lib/pq"
)

// 数据模型定义
type User struct {
    ID        int        `json:"id"`
    Name      string     `json:"name"`
    Email     string     `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// DBConfig 数据库配置
type DBConfig struct {
    Username    string
    Password    string
    Host        string
    Port        int
    DBName      string
}

// Database 封装数据库连接
type Database struct {
    *sql.DB
    config DBConfig
}

// NewDatabase 初始化数据库连接
func NewDatabase(config DBConfig) (*Database, error) {
    var db *sql.DB
    var err error

    // PostgreSQL
    // connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Username, config.Password, config.DBName)
    // db, err = sql.Open("postgres", connectionString)

    // MySQL
    mysqlConfig := mysql.Config{
        User:   config.Username,
        Passwd: config.Password,
        Net:    "tcp",
        Addr:   fmt.Sprintf("%s:%d", config.Host, config.Port),
        DBName: config.DBName,
    }
    connectionString := mysqlConfig.FormatDSN()
    db, err = sql.Open("mysql", connectionString)

    if err != nil {
        return nil, err
    }

    // 设置数据库连接池参数
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    return &Database{db, config}, nil
}

func main() {
    // 数据库配置
    config := DBConfig{
        Username:    "your_username",
        Password:    "your_password",
        Host:        "localhost",
        Port:        3306,
        DBName:      "your_db",
    }

    // 初始化数据库连接
    db, err := NewDatabase(config)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // 确保数据库连接正常
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    // 创建Echo实例
    e := echo.New()

    // 定义路由并处理HTTP请求
    e.GET("/users", func(c echo.Context) error {
        // 查询所有用户
        users, err := GetAllUsers(db)
        if err != nil {
            return err
        }

        // 返回JSON响应
        return c.JSON(200, users)
    })

    // 启动Echo服务器
    e.Logger.Fatal(e.Start(":8080"))
}

// GetAllUsers 查询所有用户
func GetAllUsers(db *Database) ([]User, error) {
    var users []User
    rows, err := db.Query("SELECT id, name, email, created_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}
