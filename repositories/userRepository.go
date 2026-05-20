package repositories

import (
	"context"
	"go-backend/config"
	"go-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	collection string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		collection: "users",
	}
}

func (r *UserRepository) CreateUser(user models.User) (interface{}, error) {
	collection := config.DB.Database("golang").Collection(r.collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"name":  user.Name,
		"email": user.Email,
	})

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *UserRepository) GetUserList() (users []models.User, err error) {
	collection := config.DB.Database("golang").Collection(r.collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	curser, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer curser.Close(ctx)

	for curser.Next(ctx) {
		var user models.User
		if err := curser.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(id primitive.ObjectID) (user models.User, err error) {
	collection := config.DB.Database("golang").Collection(r.collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, err

}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) (interface{}, error) {
	collection := config.DB.Database("golang").Collection(r.collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}
