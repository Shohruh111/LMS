package models

type CourseReport struct {
	Id            string `json:"id"`
	CourseId      string `json:"course_id"`
	PercentOfDone int    `json:"percent_of_done"`
	RemainingExam string `json:"remaining_exam"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type CourseReportCreate struct {
}

type CourseReportUpdate struct {
}

type CourseReportPrimaryKey struct {
}

type CourseReportGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CourseReportGetListResponse struct {
	Count         int             `json:"count"`
	CourseReports []*CourseReport `json:"course_reports"`
}
