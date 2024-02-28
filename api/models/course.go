package models

type Course struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Photo          string   `json:"photo"`
	ForWho         string   `json:"for_who"`
	Type           string   `json:"type"`
	WeeklyNumber   int      `json:"weekly_number"`
	Duration       string   `json:"duration"`
	Price          int      `json:"price"`
	BeginningDate  string   `json:"beginning_date_course"`
	EndDate        string   `json:"end_date"`
	NamesOfLessons []string `json:"lesson_names"`
	VideoOfLessons []string `json:"video_lessons"`
	Groups         []*Group  `json:"groups"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type CourseCreate struct {
	Name          string `json:"name"`
	Photo         string `json:"photo"`
	ForWho        string `json:"for_who"`
	Type          string `json:"type"`
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
	ForWho        string `json:"for_who"`
	Type          string `json:"type"`
	WeeklyNumber  int    `json:"weekly_number"`
	Duration      string `json:"duration"`
	Price         int    `json:"price"`
	BeginningDate string `json:"beginning_date_course"`
}

type CoursePrimaryKey struct {
	Id string `json:"id"`
}

type CourseGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type CourseOfUsers struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
}

type CourseOfUsersGetListResponse struct {
	Count         int              `json:"count"`
	CourseOfUsers []*CourseOfUsers `json:"course_of_users"`
}

type CourseGetListResponse struct {
	Count   int       `json:"count"`
	Courses []*Course `json:"courses"`
}
