package services

import (
	"go-backend/models"
	"go-backend/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {

	return &UserService{
		Repository: repository,
	}
}

func  (s *UserService) CreateUser(user models.User) (interface{}, error) {
	return s.Repository.CreateUser(user)
}

func (s *UserService) GetUserList() ([]models.User, error) {
	return s.Repository.GetUserList()
}

func (s *UserService) GetUserById(id string) (models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.User{}, err
	}
	return s.Repository.GetUserById(objectId)
}

func (s *UserService) DeleteUser(id string) (interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	
	if err != nil {
		return 0, err
	}

	return s.Repository.DeleteUser(objectId)
}