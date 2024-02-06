package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *models.UserCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "users"(id, first_name, last_name, email, phone_number)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string

		id          sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		email       sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			phone_number,
			created_at, 
			updated_at
		FROM "users"
		WHERE id = $1
	`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&email,
		&phoneNumber,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Email:       email.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (u *userRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			email,
			phone_number,
			created_at,
			updated_at
		FROM "users"
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			email       sql.NullString
			phoneNumber sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&firstName,
			&lastName,
			&email,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Email:       email.String,
			PhoneNumber: phoneNumber.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}
	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *models.UserUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"users"
		SET
			first_name = :first_name,
			last_name = :last_name,
			email = :email,
			phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.Id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"email":        req.Email,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "users" WHERE id = $1`, req.Id)
	if err != nil {
		return err
	}

	return nil
}
