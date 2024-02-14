package models

type CourseReport struct {
	Id         string `json:"id"`
	CourseId   string `json:"course_id"`
	Students   int    `json:"students"`
	Type       string `json:"type"`
	DoneAll    int    `json:"done_all"`
	NotDone    int    `json:"not_done"`
	NotStarted int    `json:"not_started"`
	Status     bool   `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CourseReportCreate struct {
	CourseId   string `json:"course_id"`
	Students   int    `json:"students"`
	Type       string `json:"type"`
	DoneAll    int    `json:"done_all"`
	NotDone    int    `json:"not_done"`
	NotStarted int    `json:"not_started"`
	Status     bool   `json:"status"`
}

type CourseReportUpdate struct {
	Id         string `json:"id"`
	CourseId   string `json:"course_id"`
	Students   int    `json:"students"`
	Type       string `json:"type"`
	DoneAll    int    `json:"done_all"`
	NotDone    int    `json:"not_done"`
	NotStarted int    `json:"not_started"`
	Status     bool   `json:"status"`
}

type CourseReportPrimaryKey struct {
	Id       string `json:"id"`
	CourseId string `json:"course_id"`
}

type CourseReportGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CourseReportGetListResponse struct {
	Count         int             `json:"count"`
	CourseReports []*CourseReport `json:"course_reports"`
}
