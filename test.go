package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "bytes"
)

// User struct to hold user data
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Global variable for the Laravel API URL
const apiURL = "http://localhost:8000/api/users"

// Function to fetch all users
func getUsers() ([]User, error) {
    resp, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var users []User
    if err := json.Unmarshal(body, &users); err != nil {
        return nil, err
    }

    return users, nil
}

// Function to create a new user
func createUser(user User) (*User, error) {
    jsonData, err := json.Marshal(user)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var createdUser User
    if err := json.Unmarshal(body, &createdUser); err != nil {
        return nil, err
    }

    return &createdUser, nil
}

//function for delete user

func main() {
    // Example: Fetch users from Laravel API

    users, err := getUsers()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Users from Laravel API: ", users)

    // Create a new user

    // newUser := User{Name: "John Doe", Email: "john.doe@example.com"}
    // createdUser, err := createUser(newUser)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println("Created user:", createdUser)

    // Delete User

 
}
