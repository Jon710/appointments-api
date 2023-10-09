package types

type Patient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Doctor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

type Appointment struct {
	ID        int    `json:"id"`
	PatientID int    `json:"patient_id"`
	DoctorID  int    `json:"doctor_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}
