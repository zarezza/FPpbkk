package models

import (
	"database/sql"
	"final-project/config"
	"final-project/entities"
	"fmt"
)

type PatientModel struct {
	conn *sql.DB
}

func NewPatientModel() *PatientModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PatientModel{
		conn: conn,
	}
}

func (p *PatientModel) FindAll() ([]entities.Patient, error) {
	rows, err := p.conn.Query("select * from patient")
	if err != nil {
		return []entities.Patient{}, err
	}
	defer rows.Close()

	var dataPatient []entities.Patient
	for rows.Next() {
		var patient entities.Patient
		rows.Scan(
			&patient.Id,
			&patient.FullName,
			&patient.SocialNumber,
			&patient.Gender,
			&patient.Birthplace,
			&patient.Birthdate,
			&patient.Address,
			&patient.PhoneNumber)

		if patient.Gender == "1" {
			patient.Gender = "Male"
		} else {
			patient.Gender = "Female"
		}

		// birth_date, _ := time.Parse("2006-01-02", patient.Birthdate)
		// patient.Birthdate = birth_date.Format("02-01-2006")

		dataPatient = append(dataPatient, patient)

	}
	return dataPatient, nil
}

func (p *PatientModel) Create(patient entities.Patient) bool {
	result, err := p.conn.Exec("insert into patient (full_name, social_number, gender, birthplace, birthdate, address, phone_number) values(?,?,?,?,?,?,?)",
		patient.FullName, patient.SocialNumber, patient.Gender, patient.Birthplace, patient.Birthdate, patient.Address, patient.PhoneNumber)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *PatientModel) Find(id int64, patient *entities.Patient) error {
	return p.conn.QueryRow("select * from patient where id = ?", id).Scan(
		&patient.Id,
		&patient.FullName,
		&patient.SocialNumber,
		&patient.Gender,
		&patient.Birthdate,
		&patient.Birthplace,
		&patient.Address,
		&patient.PhoneNumber,
	)
}

func (p *PatientModel) Update(patient entities.Patient) error {
	_, err := p.conn.Exec(
		"update patient set full_name=?, social_number=?, gender=?, birthplace=?, birthdate=?, address=?, phone_number=? where id=?",
		patient.FullName, patient.SocialNumber, patient.Gender, patient.Birthplace, patient.Birthdate, patient.Address, patient.PhoneNumber, patient.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *PatientModel) Delete(id int64) {
	p.conn.Exec("delete from patient where id = ?", id)
}
