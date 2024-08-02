package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mongodb_connect "github.com/rajpurkait9/hotel_backend/DBs/mongodb"
	Routes "github.com/rajpurkait9/hotel_backend/routes"
	Utils "github.com/rajpurkait9/hotel_backend/utils"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/health", HealthCheckHandler).Methods("GET")
	Routes.Routes(r)
	r.Use(Utils.LoggingMiddleware)
	http.Handle("/", r)

	mongodb_connect.MongoConnect()
	fmt.Println("Server is is running on port:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
