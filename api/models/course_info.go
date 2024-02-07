package models

type CourseInfo struct {
	Id            string `json:"id"`
	CourseId       string `json:"course_id"`
	PercentOfDone int    `json:"percent_of_done"`
	RemainingExam string `json:"remaining_exam"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type CourseInfoCreate struct {
	
}

type CourseInfoUpdate struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CourseInfoPrimaryKey struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CourseInfoGetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type CourseInfoGetListResponse struct {
	Count       int           `json:"count"`
	CourseInfos []*CourseInfo `json:"CourseInfos"`
}
