package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO users(id, login, password, name, age)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Login,
		req.Password,
		req.Name,
		req.Age,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string

		id       sql.NullString
		login    sql.NullString
		password sql.NullString
		name     sql.NullString
		age      int
	)

	query = `
		SELECT
			id,
			login,
			password,
			name,
			age
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&login,
		&password,
		&name,
		&age,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:       id.String,
		Login:    login.String,
		Password: password.String,
		Name:     name.String,
		Age:      age,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp   = &models.GetListUserResponse{}
		query  string
		where  = " WHERE TRUE"
		order  = "ORDER BY DESC"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			login,
			password,
			name,
			age
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + order + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id       sql.NullString
			login    sql.NullString
			password sql.NullString
			name     sql.NullString
			age      int
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&login,
			&password,
			&name,
			&age,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:       id.String,
			Login:    login.String,
			Password: password.String,
			Name:     name.String,
			Age:      age,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			users
		SET
			login = $1,
			password = $2
		WHERE id = $3
	`

	result, err := r.db.Exec(ctx, query,
		req.Id,
		req.Login,
		req.Password,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
