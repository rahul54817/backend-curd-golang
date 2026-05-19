package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend/config"
	"go-backend/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	collection := config.DB.Database("golang").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"name":  user.Name,
		"email": user.Email,
	})

	if err != nil {
		http.Error(w, "Failed to create User ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User Created",
		"id":      result.InsertedID,
	})
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	collection := config.DB.Database("golang").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	curser, err := collection.Find(ctx, bson.M{})

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
	}

	defer curser.Close(ctx)

	var users []models.User

	for curser.Next(ctx) {
		var user models.User

		curser.Decode(&user)

		users = append(users, user)

	}

	json.NewEncoder(w).Encode(users)

}

func getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	id := r.URL.Query().Get("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid user id ", http.StatusBadRequest)
		return
	}

	collection := config.DB.Database("golang").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	err = collection.FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&user)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(user)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=-Type", "application/json")

	id := r.URL.Query().Get("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid user id ", http.StatusBadRequest)
		return
	}

	collection := config.DB.Database("golang").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{
		"_id": objectId,
	})

	if err != nil {
		http.Error(w, "Faild to delete user ", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "User not found ", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"massege": "user deleted successfully",
	})
}

func main() {

	config.ConnectDb()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createUser(w, r)

		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				getUserById(w, r)
			} else {
				getUserList(w, r)
			}
		case http.MethodDelete:
			deleteUserById(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server Started")
	http.ListenAndServe(":8000", nil)
}
