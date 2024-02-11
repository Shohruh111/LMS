package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
}
type UserRepoI interface {
	Create(context.Context, *models.UserCreate) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UserUpdate) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
	CheckOTP(context.Context, *models.CheckEmail, int) (string, error)
	GetOTP(context.Context, *models.CheckCode) (string, error)
}
