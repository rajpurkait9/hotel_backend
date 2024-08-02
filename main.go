package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mongodb_connect "github.com/rajpurkait9/hotel_backend/DBs/mongodb"
	user_controller "github.com/rajpurkait9/hotel_backend/controllers/users"
	Utils "github.com/rajpurkait9/hotel_backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var err error
	_, err = mongodb_connect.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/health", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/api/v1/user", user_controller.CreateUser).Methods("POST")
	// user_routes.User_routes(r)
	r.Use(Utils.LoggingMiddleware)
	http.Handle("/", r)
	fmt.Println("Server is is running on port:8080")
	error := http.ListenAndServe(":8080", r)
	if error != nil {
		log.Fatal(err)
	}

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// package main

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}

// func main() {
// 	var err error
// 	client, err = connectMongoDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	router := mux.NewRouter()
// 	router.HandleFunc("/users", createUser).Methods("POST")
// 	router.HandleFunc("/users", getUsers).Methods("GET")
// 	router.HandleFunc("/users/{id}", getUser).Methods("GET")
// 	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
// 	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
// 	fmt.Println("Server is running on port:8080")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	_ = json.NewDecoder(r.Body).Decode(&user)
// 	user.ID = primitive.NewObjectID()
// 	collection := client.Database("testdb").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	result, _ := collection.InsertOne(ctx, user)
// 	json.NewEncoder(w).Encode(result)
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var users []User
// 	collection := client.Database("testdb").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var user User
// 		cursor.Decode(&user)
// 		users = append(users, user)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(users)
// }

// func getUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	var user User
// 	collection := client.Database("testdb").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	var user User
// 	_ = json.NewDecoder(r.Body).Decode(&user)
// 	collection := client.Database("testdb").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	_, err := collection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": id},
// 		bson.D{
// 			{"$set", bson.D{
// 				{"name", user.Name},
// 				{"email", user.Email},
// 			}},
// 		},
// 	)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	collection := client.Database("testdb").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode("User deleted")
// }
