package db

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectMySQL(dataSourceName string) error {
    var err error
    db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        return fmt.Errorf("无法连接到MySQL数据库: %v", err)
    }

    if err = db.Ping(); err != nil {
        return fmt.Errorf("无法 ping MySQL数据库: %v", err)
    }

    return nil
}

func GetDB() *sql.DB {
    return db
}