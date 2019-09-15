package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/me/artisan_full/bindings"
	"github.com/me/artisan_full/models"
	"github.com/me/artisan_full/renderings"
)

// CreateHelp - helps to create a new help instance in the db.
func CreateHelp(c echo.Context) error {
	resp := renderings.GeneralResponse{}
	helpReq := new(bindings.HelpRequest)

	if err := c.Bind(helpReq); err != nil {
		resp.Success = false
		resp.Message = "Unable to bind request for help"
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := helpReq.Validate(); err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, err)
	}

	// fmt.Printf("%+v", helpReq)

	// build help model
	help := new(models.Help)

	help.Age = helpReq.Age
	help.City = helpReq.City
	help.ContactAddress = helpReq.ContactAddress
	help.EL = helpReq.EL
	help.Email = helpReq.Email
	help.Experience = helpReq.Experience
	help.ExperienceLocations = helpReq.ExperienceLocations
	help.FirstName = helpReq.FirstName
	help.GuarantorCity = helpReq.GuarantorCity
	help.GuarantorContactAddress = helpReq.GuarantorContactAddress
	help.GuarantorEmail = helpReq.GuarantorEmail
	help.GuarantorFirstName = helpReq.GuarantorFirstName
	help.GuarantorLastName = helpReq.GuarantorLastName
	help.GuarantorPhoneNumber = helpReq.GuarantorPhoneNumber
	help.HaveTrainingCertificate = helpReq.HaveTrainingCertificate
	help.HoursOfWork = helpReq.HoursOfWork
	help.LastName = helpReq.LastName
	help.MOI = helpReq.MOI
	help.Meals = helpReq.Meals
	help.PhoneNumber = helpReq.PhoneNumber
	help.SchoolNames = helpReq.SchoolNames
	help.Service = helpReq.Service
	help.SuitableStory = helpReq.SuitableStory
	help.YearsOfExperience = helpReq.YearsOfExperience

	// save to the db
	result := models.GormDB.Create(&help)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "An error ocurred, please try again later.",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"msg": "Successful",
	})

	/*

		helpData := models.Help{}
		defer c.Request().Body.Close()
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Printf("Failed loading the request body for Help Data: %s", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"msg": "An error occured, please try again later.",
			})
		}

		err = json.Unmarshal(body, &helpData)
		if err != nil {
			log.Printf("Failed in unmarshalling for the create help : %s ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"msg": "An error occured, please try again later.",
			})
		}

		// validate
		err = helpData.Validate()
		// b, _ := json.Marshal(err)
		return c.JSON(http.StatusInternalServerError, err)

		// save to the db
		config.GormDB.Create(&helpData)
		return c.JSON(http.StatusOK, map[string]string{
			"msg": "Successful",
		})
	*/
}
