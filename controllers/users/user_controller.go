package user_controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	mongodb_connect "github.com/rajpurkait9/hotel_backend/DBs/mongodb"
	models_user "github.com/rajpurkait9/hotel_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000000; i++ {
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK get all users "))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models_user.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = primitive.NewObjectID()
	var collection = mongodb_connect.Collection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
}
