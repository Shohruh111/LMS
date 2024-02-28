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
