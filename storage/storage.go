package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	Phone() PhoneRepoI
}
type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}
type PhoneRepoI interface {
	Create(context.Context, *models.CreatePhone) (string, error)
	GetByID(context.Context, *models.PhonePrimaryKey) (*models.Phone, error)
	GetList(context.Context, *models.GetListPhoneRequest) (*models.GetListPhoneResponse, error)
	Update(context.Context, *models.UpdatePhone) (int64, error)
	Delete(context.Context, *models.PhonePrimaryKey) error
}
