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
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Email           string             `bson:"email"`
	Address         string             `bson:"address"`
	UserType        Role               `bson:"role"`
	PasswordHash    string             `bson:"passwordHash"`
	ProfileHeadline string             `bson:"profileHeadline"`
}

type Job struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title             string               `bson:"title" json:"title"`
	Description       string               `bson:"description" json:"description"`
	CompanyName       string               `bson:"company_name" json:"company_name"`
	Applicants        []primitive.ObjectID `bson:"applicants" json:"applicants"`
	TotalApplications int                  `bson:"total_applications" json:"total_applications"`
	PostedOn          time.Time            `bson:"posted_on" json:"posted_on"`
	PostedBy          primitive.ObjectID   `bson:"posted_by" json:"posted_by"`
}

type Profile struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ApplicantID       primitive.ObjectID `bson:"applicant_id"`
	ResumeFileAddress string             `bson:"resume_file_address"`
	Skills            []string           `bson:"skills"`
	Education         []Education        `bson:"education"`
	Experience        []Experience       `bson:"experience"`
	Name              string             `bson:"name"`
	Email             string             `bson:"email"`
	Phone             string             `bson:"phone"`
}

type Education struct {
	Institution  string `bson:"institution"`
	Degree       string `bson:"degree"`
	FieldOfStudy string `bson:"fieldOfStudy"`
	StartDate    string `bson:"startDate"`
	EndDate      string `bson:"endDate"`
}

type Experience struct {
	Company     string `bson:"company"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	StartDate   string `bson:"startDate"`
	EndDate     string `bson:"endDate"`
}
