package models

type VideoLessons struct {
	Id        string `json:"id"`
	FileName  string `json:"file_name"`
	PhotoData []byte `json:"video_data"`
}

type Lessons struct {
	Id          string `json:"id"`
	CourseId    string `json:"course_id"`
	Name        string `json:"name"`
	VideoLesson string `json:"video_lesson"`
	Status      bool   `json:"status"`
}

type LessonCreate struct {
	Name        string `json:"name"`
	CourseId    string `json:"course_id"`
	Status      bool   `json:"status"`
	VideoLesson string `json:"video_lesson"`
}

type LessonUpdate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	VideoLesson string `json:"video_lesson"`
}

type LessonPrimaryKey struct {
	ID string `json:"id"`
}

type LessonGetListRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	CourseId string `json:"course_id"`
}

type LessonGetListResponse struct {
	Count  int      `json:"count"`
	Lessons []*Lessons `json:"lessons"`
}
