package patientcontroller

import (
	"final-project/entities"
	"final-project/libraries"
	"final-project/models"
	"net/http"
	"strconv"
	"text/template"
)

var validation = libraries.NewValidation()

var patientModel = models.NewPatientModel()

func Index(response http.ResponseWriter, request *http.Request) {

	patient, _ := patientModel.FindAll()

	data := map[string]interface{}{
		"patient": patient,
	}

	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient
		patient.FullName = request.Form.Get("full_name")
		patient.SocialNumber = request.Form.Get("social_number")
		patient.Gender = request.Form.Get("gender")
		patient.Birthplace = request.Form.Get("birthplace")
		patient.Birthdate = request.Form.Get("birthdate")
		patient.Address = request.Form.Get("address")
		patient.PhoneNumber = request.Form.Get("phone_number")

		data := make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Data is succesfully stored"
			patientModel.Create(patient)
		}

		// data := map[string]interface{}{
		// 	"message": "Data is succesfully stored",
		// }

		temp, _ := template.ParseFiles("views/patient/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var patient entities.Patient
		patientModel.Find(id, &patient)

		data := map[string]interface{}{
			"patient": patient,
		}

		temp, err := template.ParseFiles("views/patient/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var patient entities.Patient
		patient.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		patient.FullName = request.Form.Get("full_name")
		patient.SocialNumber = request.Form.Get("social_number")
		patient.Gender = request.Form.Get("gender")
		patient.Birthplace = request.Form.Get("birthplace")
		patient.Birthdate = request.Form.Get("birthdate")
		patient.Address = request.Form.Get("address")
		patient.PhoneNumber = request.Form.Get("phone_number")

		data := make(map[string]interface{})

		vErrors := validation.Struct(patient)

		if vErrors != nil {
			data["patient"] = patient
			data["validation"] = vErrors
		} else {
			data["message"] = "Data is succesfully edited"
			patientModel.Update(patient)
		}

		// data := map[string]interface{}{
		// 	"message": "Data is succesfully stored",
		// }

		temp, _ := template.ParseFiles("views/patient/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	patientModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)
}
