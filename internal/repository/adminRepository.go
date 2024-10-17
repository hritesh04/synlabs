package repository

import (
	"context"
	"errors"
	"log"

	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminRepository struct {
	DB *mongo.Database
}

func NewAdminRepository(db *mongo.Database) ports.AdminRepository {
	return &adminRepository{
		DB: db,
	}
}

func (r *adminRepository) CreateJob(data *domain.Job) error {
	jobsCol := r.DB.Collection("jobs")
	_, err := jobsCol.InsertOne(context.TODO(), data)
	if err != nil {
		return errors.New("user creation failed")
	}
	return nil
}

func (r *adminRepository) GetJobByID(jobID primitive.ObjectID) (*domain.Job, error) {
	var job []domain.Job
	jobCol := r.DB.Collection("jobs")

	matchStage := bson.D{{"$match", bson.D{{"_id", jobID}}}}
	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "users"},
			{"localField", "applicants"},
			{"foreignField", "_id"},
			{"as", "applicant_details"},
		},
	}}

	scanner, err := jobCol.Aggregate(context.TODO(), mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := scanner.Close(context.TODO()); err != nil {
			log.Println("error closing db row scanner", "location", "getAllUsers")
		}

	}()
	if err := scanner.All(context.TODO(), &job); err != nil {
		return nil, err
	}
	return &job[0], nil
}

func (r *adminRepository) GetAllUsers() (*[]domain.User, error) {
	var users []domain.User
	userCol := r.DB.Collection("users")
	scanner, err := userCol.Find(context.TODO(), bson.D{})
	defer func() {
		if err := scanner.Close(context.TODO()); err != nil {
			log.Println("error closing db row scanner", "location", "getAllUsers")
		}

	}()
	if err != nil {
		return nil, err
	}
	if err := scanner.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *adminRepository) GetProfileByUserID(applicantID primitive.ObjectID) (*domain.Profile, error) {
	var profile domain.Profile
	profileCol := r.DB.Collection("profiles")
	if err := profileCol.FindOne(context.TODO(), bson.D{{"applicant_id", applicantID}}).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}
