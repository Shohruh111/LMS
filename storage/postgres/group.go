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

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) *groupRepo {
	return &groupRepo{
		db: db,
	}
}

func (u *groupRepo) Create(ctx context.Context, req *models.GroupCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "group"(id, name, course_id,status, end_date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Name,
		req.CourseId,
		req.Status,
		req.EndDate,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *groupRepo) GetByID(ctx context.Context, req *models.GroupPrimaryKey) (*models.Group, error) {

	var (
		query string

		id       sql.NullString
		name     sql.NullString
		courseId sql.NullString
		status   sql.NullBool
		endDate  sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			course_id,
			status,
			end_date
		FROM "group"
		WHERE id = $1
	`

	err := u.db.QueryRow(ctx, query, req.ID).Scan(
		&id,
		&name,
		&courseId,
		&status,
		&endDate,
	)

	if err != nil {
		return nil, err
	}

	return &models.Group{
		ID:       id.String,
		Name:     name.String,
		CourseId: courseId.String,
		Status:   status.Bool,
		EndDate:  endDate.String,
	}, nil
}

func (u *groupRepo) GetList(ctx context.Context, req *models.GroupGetListRequest) (*models.GroupGetListResponse, error) {

	var (
		resp   = &models.GroupGetListResponse{}
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
			end_date
		FROM "group" 

	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.CourseId) > 0 {
		where += " AND course_id = " + "'" + req.CourseId + "'"

	}

	query += where + offset + limit

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			name     sql.NullString
			courseId sql.NullString
			status   sql.NullBool
			endDate  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&courseId,
			&status,
			&endDate,
		)

		if err != nil {
			return nil, err
		}

		resp.Groups = append(resp.Groups, &models.Group{
			ID:               id.String,
			Name:             name.String,
			CourseId:         courseId.String,
			Status:           status.Bool,
			EndDate:          endDate.String,
			NumberOfStudents: 0,
			NotAll:           0,
			DoneAll:          0,
			Progress:         0,
		})
	}
	return resp, nil
}

func (u *groupRepo) Update(ctx context.Context, req *models.GroupUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"group"
		SET
			name = :name,
			status = :status,
			end_date = :end_date,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":       req.ID,
		"name":     req.Name,
		"status":   req.Status,
		"end_date": req.EndDate,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *groupRepo) Delete(ctx context.Context, req *models.GroupPrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "group" WHERE id = $1`, req.ID)
	if err != nil {
		return err
	}

	return nil
}
