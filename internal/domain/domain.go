package domain

import (
	"database/sql/driver"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Applicant Role = "applicant"
	Admin     Role = "admin"
)

func (r *Role) Scan(value string) error {
	*r = Role(value)
	return nil
}

func (r Role) Value() driver.Value {
	return string(r)
}

type User struct {
	ID              primitive.ObjectID
	Name            string  `bson:"name"`
	Email           string  `bson:"email"`
	Address         string  `bson:"address"`
	UserType        Role    `bson:"role"`
	PasswordHash    string  `bson:"passwordHash"`
	ProfileHeadline string  `bson:"profileHeadline"`
	Profile         Profile `bson:"profile"`
}

type Job struct {
	Title             string    `bson:"title"`
	Description       string    `bson:"description"`
	CompanyName       string    `bson:"company_name"`
	Applicants        []User    `bson:"applicants" gorm:"many2many:job_applicants;"`
	TotalApplications int       `bson:"total_applications"`
	PostedOn          time.Time `bson:"posted_on"`
	UserID            uint      `bson:"user_id"`
	PostedBy          User      `bson:"posted_by"`
}

type Profile struct {
	ApplicantID       primitive.ObjectID `bson:"applicant_id"`
	ResumeFileAddress string             `bson:"resume_file_address"`
	Skills            []string           `bson:"skills"`
	Education         []interface{}      `bson:"education"`
	Experience        []interface{}      `bson:"experience"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	Phone             string             `bson:"phone"`
}
