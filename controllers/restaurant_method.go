package controllers

type RestaurantStore struct {
    restaurant map[string]Restaurant
}

func NewRestaurantStore() *UserStore {
    return &UserStore{
        users: make(map[string]User),
    }
}

type Restaurant struct {
    Username   string
    Gender     string
    Phone      string
    Pincode    string
    IsLoggedin bool
}