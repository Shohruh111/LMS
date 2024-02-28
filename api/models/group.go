package models

type Group struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CourseId      string `json:"course_id"`
	Status        bool   `json:"status"`
	BeginningDate string `json:"beginning_date_course"`
	EndDate       string `json:"end_date"`
}
