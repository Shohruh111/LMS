package models

type Course struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Photo         string `json:"photo"`
	Description   string `json:"description"`
	WeeklyNumber  int    `json:"weekly_number"`
	Duration      string `json:"duration"`
	Price         int    `json:"price"`
	BeginningDate string `json:"beginning_date_course"`
	EndDate       string `json:"end_date"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type CourseCreate struct {
	Name          string `json:"name"`
	Photo         string `json:"photo"`
	Description   string `json:"description"`
	WeeklyNumber  int    `json:"weekly_number"`
	Duration      string `json:"duration"`
	Price         int    `json:"price"`
	BeginningDate string `json:"beginning_date_course"`
	EndDate       string `json:"end_date"`
}

type CourseUpdate struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Photo         string `json:"photo"`
	Description   string `json:"description"`
	WeeklyNumber  int    `json:"weekly_number"`
	Duration      string `json:"duration"`
	Price         int    `json:"price"`
	BeginningDate string `json:"beginning_date_course"`
	EndDate       string `json:"end_date"`
}

type CoursePrimaryKey struct {
	Id   string `json:"id"`
}

type CourseGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CourseGetListResponse struct {
	Count   int       `json:"count"`
	Courses []*Course `json:"courses"`
}
