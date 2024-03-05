package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Role() RoleRepoI
	Course() CourseRepoI
	Group() GroupRepoI
	Lesson() LessonRepoI
	Device() DeviceRepoI
}
type UserRepoI interface {
	Create(context.Context, *models.UserCreate) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UserUpdate) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
	CheckOTP(context.Context, *models.CheckEmail, int) (string, error)
	GetOTP(context.Context, *models.CheckCode) (string, error)
	UpdatePassword(context.Context, *models.UpdatePassword) (int64, string, error)
	GetAllUsersForExcel(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	GetAllCoursesForExcel(context.Context) (*models.CourseGetListResponse, error)
}

type RoleRepoI interface {
	Create(context.Context, *models.RoleCreate) (string, error)
	GetByID(context.Context, *models.RolePrimaryKey) (*models.Role, error)
	GetList(context.Context, *models.RoleGetListRequest) (*models.RoleGetListResponse, error)
	Update(context.Context, *models.RoleUpdate) (int64, error)
	Delete(context.Context, *models.RolePrimaryKey) error
}

type GroupRepoI interface {
	Create(context.Context, *models.GroupCreate) (string, error)
	GetByID(context.Context, *models.GroupPrimaryKey) (*models.Group, error)
	GetList(context.Context, *models.GroupGetListRequest) (*models.GroupGetListResponse, error)
	Update(context.Context, *models.GroupUpdate) (int64, error)
	Delete(context.Context, *models.GroupPrimaryKey) error
}

type LessonRepoI interface {
	Create(context.Context, *models.LessonCreate) (string, error)
	GetByID(context.Context, *models.LessonPrimaryKey) (*models.Lessons, error)
	GetList(context.Context, *models.LessonGetListRequest) (*models.LessonGetListResponse, error)
	Update(context.Context, *models.LessonUpdate) (int64, error)
	Delete(context.Context, *models.LessonPrimaryKey) error
}

type CourseRepoI interface {
	Create(context.Context, *models.CourseCreate) (string, error)
	GetByID(context.Context, *models.CoursePrimaryKey) (*models.Course, error)
	GetList(context.Context, *models.CourseGetListRequest) (*models.CourseGetListResponse, error)
	Update(context.Context, *models.CourseUpdate) (int64, error)
	Delete(context.Context, *models.CoursePrimaryKey) error
	UploadPhotos(context.Context, *models.VideoLessons) (string, error)
	GetPhotos(context.Context, *models.VideoLessons) (*models.VideoLessons, error)
	GetListCourseOfUsers(context.Context, *models.CoursePrimaryKey) (*models.CourseOfUsersGetListResponse, error)
}

type DeviceRepoI interface {
	Create(context.Context, *models.DeviceCreate) (string, error)
	GetList(context.Context, *models.DeviceGetListRequest) (*models.DeviceGetListResponse, error)
	Delete(context.Context, *models.DevicePrimaryKey) error
}
