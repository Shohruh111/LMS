package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"

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
		INSERT INTO "users"(id, role_id, first_name, last_name, email, phone_number, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	_, err = u.db.Exec(ctx, query,
		id,
		req.RoleId,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
		hashPassword,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		find  string

		id          sql.NullString
		roleId      sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		email       sql.NullString
		phoneNumber sql.NullString
		password    sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString

		userType sql.NullString
		where    string = " WHERE "
	)
	if len(req.Email) > 0 {
		where += " u.email = $1 "
		find = req.Email
	} else {
		where += " u.id = $1 "
		find = req.Id
	}

	query = `
		SELECT 
			u.id,
			u.role_id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone_number,
			u.password,
			u.created_at, 
			u.updated_at,

			r.type
		FROM "users" AS u
		JOIN "roles" AS r ON u.role_id = r.id
	` + where

	err := u.db.QueryRow(ctx, query, find).Scan(
		&id,
		&roleId,
		&firstName,
		&lastName,
		&email,
		&phoneNumber,
		&password,
		&createdAt,
		&updatedAt,
		&userType,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:          id.String,
		RoleId:      roleId.String,
		FirstName:   firstName.String,
		LastName:    lastName.String,
		Email:       email.String,
		PhoneNumber: phoneNumber.String,
		Password:    "",
		UserType:    userType.String,
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
			u.id,
			u.role_id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone_number,
			u.password,
			u.created_at,
			u.updated_at,

			r.type
		FROM "users" AS u
		JOIN "roles" AS r ON u.role_id = r.id
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
			roleId      sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			email       sql.NullString
			phoneNumber sql.NullString
			password    sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
			userType    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&roleId,
			&firstName,
			&lastName,
			&email,
			&phoneNumber,
			&password,
			&createdAt,
			&updatedAt,
			&userType,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			RoleId:      roleId.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Email:       email.String,
			PhoneNumber: phoneNumber.String,
			Password:    "",
			UserType:    userType.String,
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
			password = :password,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.Id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"email":        req.Email,
		"phone_number": req.PhoneNumber,
		"password":     req.Password,
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

func (u *userRepo) CheckOTP(ctx context.Context, req *models.CheckEmail, verifyCode int) (string, error) {

	var (
		requestId = uuid.New().String()
		query     string
	)

	query = `
		INSERT INTO "check_email"(request_id, email, verify_code)
		VALUES ($1, $2, $3)
	`

	_, err := u.db.Exec(ctx, query,
		requestId,
		req.Email,
		strconv.Itoa(verifyCode),
	)
	if err != nil {
		return "", err
	}

	query = `
		UPDATE "check_email"
		SET expired_at = created_at + INTERVAL '1 minute'
	`
	_, err = u.db.Exec(ctx, query)
	if err != nil {
		return "", err
	}

	return requestId, nil
}

func (u *userRepo) GetOTP(ctx context.Context, req *models.CheckCode) (string, error) {

	var (
		query string

		requestID  sql.NullString
		email      sql.NullString
		verifyCode sql.NullString
		createdAt  sql.NullString
		expiredAt  time.Time
	)

	query = `
		SELECT 
			request_id,
			email,
			verify_code,
			created_at,
			expired_at 
		FROM "check_email"
		WHERE request_id = $1
	`
	err := u.db.QueryRow(ctx, query, req.RequestID).Scan(
		&requestID,
		&email,
		&verifyCode,
		&createdAt,
		&expiredAt,
	)
	if err != nil {
		return "", err
	}

	if time.Now().After(expiredAt) {
		return "Code Expired!", errors.New("Code Expired!")
	}

	code, err := strconv.Atoi(verifyCode.String)
	cameCode, err := strconv.Atoi(req.Code)

	if code != cameCode || err != nil {
		return "InValid Code", errors.New("Invalid Code!")
	}

	return "Valid Code", nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, req *models.UpdatePassword) (int64, string, error) {
	var (
		query string
	)

	query = `
		UPDATE "users"
		SET password = $1, updated_at = NOW()
		WHERE email = $2
	`
	result, err := u.db.Exec(ctx, query, req.Password, req.Email)
	if err != nil {
		return 0, req.Email, err
	}

	return result.RowsAffected(), req.Email, nil
}
