package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type courseRepo struct {
	db *pgxpool.Pool
}

func NewCourseRepo(db *pgxpool.Pool) *courseRepo {
	return &courseRepo{
		db: db,
	}
}

func (u *courseRepo) Create(ctx context.Context, req *models.CourseCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "courses"(id, name, photo, for_who, type, weekly_number, duration, price, beginning_date_course, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Name,
		req.Photo,
		req.ForWho,
		req.Type,
		req.WeeklyNumber,
		req.Duration,
		req.Price,
		req.BeginingDate,
		req.EndDate,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *courseRepo) GetByID(ctx context.Context, req *models.CoursePrimaryKey) (*models.Course, error) {

	var (
		query string

		id            sql.NullString
		name          sql.NullString
		photo         sql.NullString
		forWho        sql.NullString
		tipe          sql.NullString
		weeklyNumber  int
		duration      sql.NullString
		price         int
		beginningDate sql.NullString
		createdAt     sql.NullString
		updatedAt     sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			photo,
			for_who,
			type,
			weekly_number,
			duration,
			price,
			beginning_date_course,
			created_at,
			updated_at
		FROM "courses" 
		WHERE id = $1
	`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&photo,
		&forWho,
		&tipe,
		&weeklyNumber,
		&duration,
		&price,
		&beginningDate,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Course{
		Id:            id.String,
		Name:          name.String,
		Photo:         photo.String,
		ForWho:        forWho.String,
		Type:          tipe.String,
		WeeklyNumber:  weeklyNumber,
		Duration:      duration.String,
		Price:         price,
		BeginningDate: beginningDate.String,
		CreatedAt:     createdAt.String,
		UpdatedAt:     updatedAt.String,
	}, nil
}

func (u *courseRepo) GetList(ctx context.Context, req *models.CourseGetListRequest) (*models.CourseGetListResponse, error) {

	var (
		resp   = &models.CourseGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0 "
		limit  = " LIMIT 10 "
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			photo,
			for_who,
			type,
			weekly_number,
			duration,
			price,
			beginning_date_course,
			created_at,
			updated_at
		FROM "courses" 
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
			id            sql.NullString
			name          sql.NullString
			photo         sql.NullString
			forWho        sql.NullString
			tipe          sql.NullString
			weeklyNumber  int
			duration      sql.NullString
			price         int
			beginningDate sql.NullString
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&photo,
			&forWho,
			&tipe,
			&weeklyNumber,
			&duration,
			&price,
			&beginningDate,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Courses = append(resp.Courses, &models.Course{
			Id:            id.String,
			Name:          name.String,
			Photo:         photo.String,
			ForWho:        forWho.String,
			Type:          tipe.String,
			WeeklyNumber:  weeklyNumber,
			Duration:      duration.String,
			Price:         price,
			BeginningDate: beginningDate.String,
			CreatedAt:     createdAt.String,
			UpdatedAt:     updatedAt.String,
		})
	}
	return resp, nil
}

func (u *courseRepo) Update(ctx context.Context, req *models.CourseUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"courses"
		SET
			name = :name,
			photo = :photo,
			for_who = :for_who,
			type = :type,
			weekly_number = :weekly_number,
			duration = :duration,
			price = :price,
			beginning_date_course = :beginning_date_course,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":                    req.Id,
		"name":                  req.Name,
		"photo":                 req.Photo,
		"for_who":               req.ForWho,
		"type":                  req.Type,
		"weekly_number":         req.WeeklyNumber,
		"duration":              req.Duration,
		"price":                 req.Price,
		"beginning_date_course": req.BeginningDate,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *courseRepo) Delete(ctx context.Context, req *models.CoursePrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "courses" WHERE id = $1`, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseRepo) GetListCourseOfUsers(ctx context.Context, req *models.CoursePrimaryKey) (*models.CourseOfUsersGetListResponse, error) {

	var (
		resp  = &models.CourseOfUsersGetListResponse{}
		query string
		where = " WHERE id = '" + req.Id + "'"
	)

	query = `
		SELECT 
			COUNT(*) OVER(),
			u.phone_number,
			u.first_name,
			u.last_name,
			u.email,

			cu.created_at

		FROM "course" as c
		JOIN "course_of_users" AS cu ON c.id = cu.course_id
		JOIN "users" AS u ON cu.user_id = u.id
	`

	query += where

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			phoneNumber sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			email       sql.NullString
			createdAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&phoneNumber,
			&firstName,
			&lastName,
			&email,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}

		resp.CourseOfUsers = append(resp.CourseOfUsers, &models.CourseOfUsers{
			PhoneNumber: phoneNumber.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Email:       email.String,
			CreatedAt:   createdAt.String,
		})

	}

	return resp, nil
}
