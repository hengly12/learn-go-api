package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Alice", },
    {ID: 2, Name: "Bob", },
}

var dummyUsers = map[string]string{
    "admin": "password",
    "user1": "123456",
    "test":  "test123",
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Message string `json:"message"`
    Token   string `json:"token,omitempty"` 
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    for _, user := range users {
        if user.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(user)
            return
        }
    }

    http.Error(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    json.NewDecoder(r.Body).Decode(&newUser)
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }
    
    var updatedUser User
    json.NewDecoder(r.Body).Decode(&updatedUser)
    for i, user := range users {
        if user.ID == id {
            users[i].Name = updatedUser.Name
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(users[i])
            return
        }
    }
    
    http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "User not found", http.StatusNotFound)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var creds LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    password, exists := dummyUsers[creds.Username]
    if !exists || password != creds.Password {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    response := LoginResponse{
        Message: "Login successful",
        Token:   "dummy-token-123456",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}