package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	DB *mongo.Database
}

func NewUserRepository(db *mongo.Database) ports.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(data *domain.User) error {
	userCol := r.DB.Collection("users")
	_, err := userCol.InsertOne(context.TODO(), data)
	fmt.Println("here")
	if err != nil {
		fmt.Println(err)
		return errors.New("user creation failed")
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	userCol := r.DB.Collection("users")
	err := userCol.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("No user found with the email")
	}
	return &user, nil
}

func (r *userRepository) GetAllJobs() {

}

func (r *userRepository) AddUserToJob() {

}

func (r *userRepository) CreateProfile(data *domain.Profile) error {
	userCol := r.DB.Collection("profiles")
	_, err := userCol.InsertOne(context.TODO(), data)
	if err != nil {
		return errors.New("user creation failed")
	}
	return nil
}
