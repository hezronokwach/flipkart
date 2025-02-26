package main

import (
    "fmt"
    "flipkart/controllers"
)

func main() {
    store := controllers.NewUserStore()

    err := store.RegisterUser("JohnDoe", "Male", "1234567890", 123456)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    err = store.RegisterUser("JaneDoe", "Female", "9876543210", 654321)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    users := store.ListUsers()
    fmt.Println("All users:", users)
}