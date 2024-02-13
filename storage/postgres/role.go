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

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) *roleRepo {
	return &roleRepo{
		db: db,
	}
}

func (u *roleRepo) Create(ctx context.Context, req *models.RoleCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "roles"(id, type)
		VALUES ($1, $2)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Type,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *roleRepo) GetByID(ctx context.Context, req *models.RolePrimaryKey) (*models.Role, error) {

	var (
		query string
		find  string
		where string = " WHERE "

		id        sql.NullString
		typeRole  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	if len(req.Id) > 0 {
		where += " id = $1 "
		find = req.Id
	} else if len(req.Type) > 0 {
		where += "type = $1"
		find = req.Type
	}

	query = `
		SELECT 
			id,
			type,
			created_at,
			updated_at
		FROM "roles"
	` + where

	err := u.db.QueryRow(ctx, query, find).Scan(
		&id,
		&typeRole,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Role{
		Id:        id.String,
		Type:      typeRole.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (u *roleRepo) GetList(ctx context.Context, req *models.RoleGetListRequest) (*models.RoleGetListResponse, error) {

	var (
		resp   = &models.RoleGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			type,
			created_at,
			updated_at
		FROM "roles" 
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
			id        sql.NullString
			typeRole  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&typeRole,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Roles = append(resp.Roles, &models.Role{
			Id:        id.String,
			Type:      typeRole.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}
	return resp, nil
}

func (u *roleRepo) Update(ctx context.Context, req *models.RoleUpdate) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"roles"
		SET
			type = :type,
			updated_at = NOW()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":   req.Id,
		"type": req.Type,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *roleRepo) Delete(ctx context.Context, req *models.RolePrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "roles" WHERE id = $1`, req.Id)
	if err != nil {
		return err
	}

	return nil
}
