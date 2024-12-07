package main

import (
	patientcontroller "final-project/controllers/PatientController"
	"net/http"
)

func main() {
	http.HandleFunc("/", patientcontroller.Index)
	http.HandleFunc("/patient", patientcontroller.Index)
	http.HandleFunc("/patient/index", patientcontroller.Index)
	http.HandleFunc("/patient/add", patientcontroller.Add)
	http.HandleFunc("/patient/edit", patientcontroller.Edit)
	http.HandleFunc("/patient/delete", patientcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
