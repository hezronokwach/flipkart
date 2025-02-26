package controllers

import (
    "errors"
    "strconv"
)

type UserStore struct {
    users map[string]User
}

func NewUserStore() *UserStore {
    return &UserStore{
        users: make(map[string]User),
    }
}

type User struct {
    Username   string
    Gender     string
    Phone      string
    Pincode    string
    IsLoggedin bool
}

func (store *UserStore) RegisterUser(username, gender, phone string, pin int) error {
    if username == "" {
        return errors.New("username cannot be empty")
    }
    if gender != "Male" && gender != "Female" {
        return errors.New("invalid gender")
    }
    if len(phone) != 10 {
        return errors.New("invalid phone number")
    }
    pincode := strconv.Itoa(pin)
    if len(pincode) != 6 {
        return errors.New("invalid pincode")
    }

    // Check if user exists
    if _, exists := store.users[phone]; exists {
        return errors.New("phone number already exists")
    }

    // Store new user
    store.users[phone] = User{
        Username:   username,
        Gender:     gender,
        Phone:      phone,
        Pincode:    pincode,
        IsLoggedin: false,
    }

    return nil
}

// func (store *UserStore) GetUser(username string) (User, error) {
//     if user, exists := store.users[username]; exists {
//         return user, nil
//     }
//     return User{}, errors.New("user not found")
// }

func (store *UserStore) ListUsers() []User {
    users := make([]User, 0, len(store.users))
    for _, user := range store.users {
        users = append(users, user)
    }
    return users
}

