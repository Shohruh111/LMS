package postgres

// import (
// 	"app/api/models"
// 	"app/pkg/helper"
// 	"context"
// 	"database/sql"
// 	"fmt"

// 	uuid "github.com/google/uuid"
// 	"github.com/jackc/pgx/v4/pgxpool"
// )

// type courseReportRepo struct {
// 	db *pgxpool.Pool
// }

// func NewCourseReportRepo(db *pgxpool.Pool) *courseReportRepo {
// 	return &courseReportRepo{
// 		db: db,
// 	}
// }

// func (u *courseReportRepo) Create(ctx context.Context, req *models.CourseReportCreate) (string, error) {

// 	var (
// 		id    = uuid.New().String()
// 		query string
// 	)

// 	query = `
// 		INSERT INTO "course_report"(id, course_id, students, type, done_all, not_done, not_started, status)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
// 	`

// 	_, err := u.db.Exec(ctx, query,
// 		id,
// 		req.CourseId,
// 		req.Students,
// 		req.Type,
// 		req.DoneAll,
// 		req.NotDone,
// 		req.NotStarted,
// 		req.Status,
// 	)

// 	if err != nil {
// 		return "", err
// 	}

// 	return id, nil
// }

// func (u *courseReportRepo) GetByID(ctx context.Context, req *models.CourseReportPrimaryKey) (*models.CourseReport, error) {

// 	var (
// 		query string
// 		find  string

// 		id         sql.NullString
// 		courseId   sql.NullString
// 		students   sql.NullString
// 		typeof     sql.NullString
// 		doneAll    sql.NullString
// 		notDone    sql.NullString
// 		notStarted sql.NullString
// 		status     bool
// 		createdAt  sql.NullString
// 		updatedAt  sql.NullString

// 		where    string = " WHERE "
// 	)
// 	if len(req.Email) > 0 {
// 		where += " u.email = $1 "
// 		find = req.Email
// 	} else {
// 		where += " u.id = $1 "
// 		find = req.Id
// 	}

// 	query = `
// 		SELECT 
// 			u.id,
// 			u.role_id,
// 			u.first_name,
// 			u.last_name,
// 			u.email,
// 			u.phone_number,
// 			u.password,
// 			u.created_at, 
// 			u.updated_at,

// 			r.type
// 		FROM "users" AS u
// 		JOIN "roles" AS r ON u.role_id = r.id
// 	` + where

// 	err := u.db.QueryRow(ctx, query, find).Scan(
// 		&id,
// 		&roleId,
// 		&firstName,
// 		&lastName,
// 		&email,
// 		&phoneNumber,
// 		&password,
// 		&createdAt,
// 		&updatedAt,
// 		&userType,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.User{
// 		Id:          id.String,
// 		RoleId:      roleId.String,
// 		FirstName:   firstName.String,
// 		LastName:    lastName.String,
// 		Email:       email.String,
// 		PhoneNumber: phoneNumber.String,
// 		Password:    password.String,
// 		UserType:    userType.String,
// 		CreatedAt:   createdAt.String,
// 		UpdatedAt:   updatedAt.String,
// 	}, nil
// }

// func (u *courseReportRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

// 	var (
// 		resp   = &models.UserGetListResponse{}
// 		query  string
// 		where  = " WHERE TRUE"
// 		offset = " OFFSET 0"
// 		limit  = " LIMIT 10"
// 	)

// 	query = `
// 		SELECT
// 			COUNT(*) OVER(),
// 			id,
// 			role_id,
// 			first_name,
// 			last_name,
// 			email,
// 			phone_number,
// 			password,
// 			created_at,
// 			updated_at
// 		FROM "users" 
// 	`

// 	if req.Offset > 0 {
// 		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
// 	}

// 	if req.Limit > 0 {
// 		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
// 	}

// 	query += where + offset + limit

// 	rows, err := u.db.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var (
// 			id          sql.NullString
// 			roleId      sql.NullString
// 			firstName   sql.NullString
// 			lastName    sql.NullString
// 			email       sql.NullString
// 			phoneNumber sql.NullString
// 			password    sql.NullString
// 			createdAt   sql.NullString
// 			updatedAt   sql.NullString
// 		)

// 		err := rows.Scan(
// 			&resp.Count,
// 			&id,
// 			&roleId,
// 			&firstName,
// 			&lastName,
// 			&email,
// 			&phoneNumber,
// 			&password,
// 			&createdAt,
// 			&updatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		resp.Users = append(resp.Users, &models.User{
// 			Id:          id.String,
// 			RoleId:      roleId.String,
// 			FirstName:   firstName.String,
// 			LastName:    lastName.String,
// 			Email:       email.String,
// 			PhoneNumber: phoneNumber.String,
// 			Password:    password.String,
// 			CreatedAt:   createdAt.String,
// 			UpdatedAt:   updatedAt.String,
// 		})
// 	}
// 	return resp, nil
// }

// func (u *courseReportRepo) Update(ctx context.Context, req *models.UserUpdate) (int64, error) {

// 	var (
// 		query  string
// 		params map[string]interface{}
// 	)

// 	query = `
// 		UPDATE
// 			"users"
// 		SET
// 			first_name = :first_name,
// 			last_name = :last_name,
// 			email = :email,
// 			phone_number = :phone_number,
// 			updated_at = NOW()
// 		WHERE id = :id
// 	`
// 	params = map[string]interface{}{
// 		"id":           req.Id,
// 		"first_name":   req.FirstName,
// 		"last_name":    req.LastName,
// 		"email":        req.Email,
// 		"phone_number": req.PhoneNumber,
// 	}

// 	query, args := helper.ReplaceQueryParams(query, params)
// 	result, err := u.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.RowsAffected(), nil
// }

// func (u *courseReportRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

// 	_, err := u.db.Exec(ctx, `DELETE FROM "users" WHERE id = $1`, req.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
