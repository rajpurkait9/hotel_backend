package user_routes

import (
	"net/http"

	"github.com/gorilla/mux"
	user_controller "github.com/rajpurkait9/hotel_backend/controllers/users"
)

func User_routes(r *mux.Router) {
	r.HandleFunc("/api/v1/all", user_controller.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/create", user_controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/update", user_controller.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("api/v1/delete", user_controller.DeleteUser).Methods(http.MethodDelete)
}
