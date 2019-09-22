package models

import (
	"github.com/jinzhu/gorm"
)

// Help - Model for help table
type Help struct {
	gorm.Model
	FirstName      string `gorm:"type:varchar(50); not null" json:"first_name"`
	LastName       string `gorm:"type:varchar(50); not null" json:"last_name"`
	Age            int    `gorm:"size:3; not null" json:"age"`
	PhoneNumber    string `gorm:"type:varchar(20); not null" json:"phone_number"`
	Email          string `gorm:"type:varchar(100);unique_index; not null" json:"email"`
	ContactAddress string `gorm:"type:text" json:"contact_address"`
	City           string `gorm:"type:varchar(100)" json:"city"`
	MOI            string `gorm:"type:char(15)" json:"means_of_identification"` // means_of_identification
	EL             string `gorm:"type:char(15)" json:"education_level"`         // education_level

	// job details
	Service                 string `gorm:"type:varchar(50);not null" json:"service"`
	Meals                   string `gorm:"type:text" json:"meals"` // can be null
	Experience              bool   `gorm:"DEFAULT:false" json:"experience"`
	YearsOfExperience       int    `gorm:"size:3;default:0" json:"years_of_experience"`
	ExperienceLocations     string `gorm:"type:text" json:"experience_locations"` // seperated by commas
	HaveTrainingCertificate bool   `gorm"default:false" json:"have_training_experience"`
	SchoolNames             string `gorm:"type:text" json:"school_names"`
	SuitableStory           string `gorm:"type:text; not null" json:"suitable_story"`
	HoursOfWork             int    `gorm:"size:5; not null" json:"hours_of_work"`

	// guarantor
	GuarantorFirstName      string `gorm:"type:varchar(50); not null" json:"guarantor_first_name"`
	GuarantorLastName       string `gorm:"type:varchar(50); not null" json:"guarantor_lastname"`
	GuarantorPhoneNumber    string `gorm:"type:varchar(20); not null" json:"guarantor_phone_number"`
	GuarantorEmail          string `gorm:"type:varchar(100); not null" json:"guarantor_email"`
	GuarantorContactAddress string `gorm:"type:text; not null" json:"guarantor_contact_address"`
	GuarantorCity           string `gorm:"type:varchar(100); not null" json:"guarantor_city"`
	GuarantorMOI            string `gorm:"type:char(15); default:''" json:"guarantor_moi"`
}
