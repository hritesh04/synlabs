package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *userRepository) CheckUserExists(userID primitive.ObjectID) error {
	userCol := r.DB.Collection("users")
	if err := userCol.FindOne(context.TODO(), bson.D{{"_id", userID}}).Err(); err == mongo.ErrNoDocuments {
		return errors.New("No user found with the email")
	}
	return nil
}

func (r *userRepository) GetAllJobs() (*[]domain.Job, error) {
	var jobs []domain.Job
	jobCol := r.DB.Collection("jobs")

	scanner, err := jobCol.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err := scanner.All(context.TODO(), &jobs); err != nil {
		return nil, err
	}
	return &jobs, nil
}

func (r *userRepository) AddUserToJob(userID, jobID primitive.ObjectID) error {
	jobCol := r.DB.Collection("jobs")
	update := bson.M{
		"$inc":  bson.M{"total_applicants": 1},
		"$push": bson.M{"applicants": userID},
	}
	if err := jobCol.FindOneAndUpdate(context.TODO(), bson.D{{"_id", jobID}}, update); err != nil {
		return err.Err()
	}
	return nil
}

func (r *userRepository) CreateProfile(data *domain.Profile) error {
	userCol := r.DB.Collection("profiles")
	_, err := userCol.InsertOne(context.TODO(), data)
	if err != nil {
		return errors.New("user creation failed")
	}
	return nil
}
