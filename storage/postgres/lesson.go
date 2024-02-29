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

type lessonRepo struct {
	db *pgxpool.Pool
}

func NewLessonRepo(db *pgxpool.Pool) *lessonRepo {
	return &lessonRepo{
		db: db,
	}
}

func (u *lessonRepo) Create(ctx context.Context, req *models.LessonCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "lessons"(id, name, course_id,status, video_lesson)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Name,
		req.CourseId,
		req.Status,
		req.VideoLesson,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *lessonRepo) GetByID(ctx context.Context, req *models.LessonPrimaryKey) (*models.Lessons, error) {

	var (
		query string

		id          sql.NullString
		name        sql.NullString
		courseId    sql.NullString
		status      sql.NullBool
		videoLesson sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			course_id,
			status,
			video_lesson
		FROM "lessons"
		WHERE id = $1
	`

	err := u.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&name,
		&courseId,
		&status,
		&videoLesson,
	)

	if err != nil {
		return nil, err
	}

	return &models.Lessons{
		Id:          id.String,
		Name:        name.String,
		CourseId:    courseId.String,
		Status:      status.Bool,
		VideoLesson: videoLesson.String,
	}, nil
}

func (u *lessonRepo) GetList(ctx context.Context, req *models.LessonGetListRequest) (*models.LessonGetListResponse, error) {

	var (
		resp   = &models.LessonGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			course_id,
			status,
			video_lesson
		FROM "lessons" 
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.CourseId) > 0 {
		where += " AND course_id = $1"
	}

	query += where + offset + limit

	rows, err := u.db.Query(ctx, query, req.CourseId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			name        sql.NullString
			courseId    sql.NullString
			status      sql.NullBool
			videoLesson sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&courseId,
			&status,
			&videoLesson,
		)

		if err != nil {
			return nil, err
		}

		resp.Lessons = append(resp.Lessons, &models.Lessons{
			Id:          id.String,
			Name:        name.String,
			CourseId:    courseId.String,
			Status:      status.Bool,
			VideoLesson: videoLesson.String,
		})
	}
	return resp, nil
}

func (u *lessonRepo) Update(ctx context.Context, req *models.LessonUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"lessons"
		SET
			name = :name,
			status = :status,
			video_lesson = :video_lesson,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":           req.ID,
		"name":         req.Name,
		"status":       req.Status,
		"video_lesson": req.VideoLesson,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *lessonRepo) Delete(ctx context.Context, req *models.LessonPrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "lessons" WHERE id = $1`, req.ID)
	if err != nil {
		return err
	}

	return nil
}
