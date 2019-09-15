package bindings

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// HelpRequest - this is the format an help would be sent to this app
type HelpRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Age            int    `json:"age"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	ContactAddress string `json:"contact_address"`
	City           string `json:"city"`
	MOI            string `json:"means_of_identification"` // means_of_identification
	EL             string `json:"education_level"`         // education_level

	// job details
	Service                 string `json:"service"`
	Meals                   string `json:"meals"` // can be null
	Experience              bool   `json:"experience"`
	YearsOfExperience       int    `json:"years_of_experience"`
	ExperienceLocations     string `json:"experience_locations"` // seperated by commas
	HaveTrainingCertificate bool   `json:"have_training_experience"`
	SchoolNames             string `json:"school_names"`
	SuitableStory           string `json:"suitable_story"`
	HoursOfWork             int    `json:"hours_of_work"`

	// guarantor
	GuarantorFirstName      string `json:"guarantor_first_name"`
	GuarantorLastName       string `json:"guarantor_last_name"`
	GuarantorPhoneNumber    string `json:"guarantor_phone_number"`
	GuarantorEmail          string `json:"guarantor_email"`
	GuarantorContactAddress string `json:"guarantor_contact_address"`
	GuarantorCity           string `json:"guarantor_city"`
}

// Validate - this is the validator class for the request.
func (h HelpRequest) Validate() error {
	return validation.ValidateStruct(&h,
		validation.Field(&h.FirstName, validation.Required, validation.Length(5, 50)),
		validation.Field(&h.LastName, validation.Required, validation.Length(5, 50)),
		validation.Field(&h.Age, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&h.Email, validation.Required, is.Email),
		validation.Field(&h.PhoneNumber, validation.Required, validation.Length(11, 11)),
		validation.Field(&h.ContactAddress, validation.Required, validation.Length(5, 500)),
		validation.Field(&h.City, validation.Required, validation.Length(2, 50)),
		validation.Field(&h.MOI, validation.Required),
		validation.Field(&h.EL, validation.Required),
		validation.Field(&h.Service, validation.Required),
		validation.Field(&h.Meals, validation.Length(0, 500)),

		// job details
		validation.Field(&h.SuitableStory, validation.Required),
		validation.Field(&h.HoursOfWork, validation.Required, validation.Max(120)),

		// guarantor
		validation.Field(&h.GuarantorFirstName, validation.Required, validation.Length(5, 50)),
		validation.Field(&h.GuarantorLastName, validation.Required, validation.Length(5, 50)),
		validation.Field(&h.GuarantorEmail, validation.Required, is.Email),
		validation.Field(&h.GuarantorPhoneNumber, validation.Required, validation.Length(11, 11)),
		validation.Field(&h.GuarantorContactAddress, validation.Required, validation.Length(5, 500)),
		validation.Field(&h.GuarantorCity, validation.Required, validation.Length(2, 50)),
	)
}
