package user_controller

import (
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000000; i++ {
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK get all users "))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user created successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
}
