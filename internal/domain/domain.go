package domain

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
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
	gorm.Model
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Address         string  `json:"address"`
	UserType        Role    `json:"role" gorm:"type:role;default:'applicant'"`
	PasswordHash    string  `json:"-"`
	ProfileHeadline string  `json:"profileHeadline"`
	Profile         Profile `json:"profile" gorm:"foreignKey:ApplicantID"`
}

type Profile struct {
	gorm.Model
	ApplicantID       uint   `json:"applicant_id"`
	ResumeFileAddress string `json:"resume_file_address"`
	Skills            string `json:"skills"`
	Education         string `json:"education"`
	Experience        string `json:"experience"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
}

type Job struct {
	gorm.Model
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	CompanyName       string    `json:"company_name"`
	Applicants        []User    `json:"applicants" gorm:"many2many:job_applicants;"`
	TotalApplications int       `json:"total_applications"`
	PostedOn          time.Time `json:"posted_on"`
	UserID            uint      `json:"user_id"`
	PostedBy          User      `json:"posted_by" gorm:"foreignkey:UserID"`
}
