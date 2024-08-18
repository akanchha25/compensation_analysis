// models/compensation.go
package models

type Compensation struct {
    Age              int     `json:"age"`
    Industry         string  `json:"industry"`
    JobTitle         string  `json:"job_title"`
    BaseSalary       float64 `json:"base_salary"`
    Currency         string  `json:"currency"`
    Location         string  `json:"location"`
    YearsOfExperience float64 `json:"years_of_experience"`
    Timestamp        string  `json:"timestamp"`
}
