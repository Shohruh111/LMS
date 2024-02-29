package models

type Group struct {
	ID       string `json:"id"`
	CourseId string `json:"course_id"`
	Name     string `json:"name"`
	Status   bool   `json:"status"`
	EndDate  string `json:"end_date"`
}
type GroupCreate struct {
	Name     string `json:"name"`
	CourseId string `json:"course_id"`
	Status   bool   `json:"status"`
	EndDate  string `json:"end_date"`
}

type GroupUpdate struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  bool   `json:"status"`
	EndDate string `json:"end_date"`
}

type GroupPrimaryKey struct {
	ID string `json:"id"`
}

type GroupGetListRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	CourseId string `json:"course_id"`
}

type GroupGetListResponse struct {
	Count  int      `json:"count"`
	Groups []*Group `json:"groups"`
}
