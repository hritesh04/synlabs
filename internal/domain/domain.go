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
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Email           string             `bson:"email" json:"email"`
	Address         string             `bson:"address" json:"address"`
	UserType        Role               `bson:"role" json:"role"`
	PasswordHash    string             `bson:"passwordHash" json:"-"`
	ProfileHeadline string             `bson:"profileHeadline" json:"profile_headline"`
}

type Job struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title             string               `bson:"title" json:"title"`
	Description       string               `bson:"description" json:"description"`
	CompanyName       string               `bson:"company_name" json:"company_name"`
	Applicants        []primitive.ObjectID `bson:"applicants" json:"-"`
	ApplicantDetails  []User               `bson:"applicant_details,omitempty" json:"applicant_details"`
	TotalApplications int                  `bson:"total_applications" json:"total_applications"`
	PostedOn          time.Time            `bson:"posted_on" json:"posted_on"`
	PostedBy          primitive.ObjectID   `bson:"posted_by" json:"posted_by"`
}

type Profile struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ApplicantID       primitive.ObjectID `bson:"applicant_id" json:"applicant_id"`
	ResumeFileAddress string             `bson:"resume_file_address" json:"resume_file_address"`
	Skills            string             `bson:"skills" json:"skills"`
	Education         string             `bson:"education" json:"education"`
	Experience        string             `bson:"experience" json:"experience"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bson:"email" json:"email"`
	Phone             string             `bson:"phone" json:"phone"`
}

// type Education struct {
// 	Institution  string `bson:"institution"`
// 	Degree       string `bson:"degree"`
// 	FieldOfStudy string `bson:"fieldOfStudy"`
// 	StartDate    string `bson:"startDate"`
// 	EndDate      string `bson:"endDate"`
// }

// type Experience struct {
// 	Company     string `bson:"company"`
// 	Title       string `bson:"title"`
// 	Description string `bson:"description"`
// 	StartDate   string `bson:"startDate"`
// 	EndDate     string `bson:"endDate"`
// }
