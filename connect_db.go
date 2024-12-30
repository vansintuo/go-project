package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    // Add other fields as necessary
}

var db *sql.DB

func initDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/vue_db"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Database connection established")
}

func getUsersFromDB() ([]User, error) {
    query := "SELECT id, name, email FROM users"
    log.Println("Executing query:", query)
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    log.Println("Query executed successfully, fetched users:", users)
    return users, nil
}

func main() {
    initDB()
    defer db.Close()

    users, err := getUsersFromDB()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Users from MySQL database: ", users)
}