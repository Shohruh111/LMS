package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
)

type phoneRepo struct {
	db *pgxpool.Pool
}

func NewPhoneRepo(db *pgxpool.Pool) *phoneRepo {
	return &phoneRepo{
		db: db,
	}
}

func (r *phoneRepo) Create(ctx context.Context, req *models.CreatePhone) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO phones(id,user_id, phone, descriprion, is_fax)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.UserId,
		req.Phone,
		req.Description,
		req.IsFax,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *phoneRepo) GetByID(ctx context.Context, req *models.PhonePrimaryKey) (*models.Phone, error) {

	var (
		query string

		user_id     sql.NullString
		phone       sql.NullString
		descriprion sql.NullString
		is_fax      bool
	)

	query = `
		SELECT
			user_id,
			phone,
			descriprion,
			is_fax
		FROM phones
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&user_id,
		&phone,
		&descriprion,
		&is_fax,
	)

	if err != nil {
		return nil, err
	}

	return &models.Phone{
		Id:          req.Id,
		UserId:      user_id.String,
		Phone:       phone.String,
		Description: descriprion.String,
		IsFax:       is_fax,
	}, nil
}

func (r *phoneRepo) GetList(ctx context.Context, req *models.GetListPhoneRequest) (*models.GetListPhoneResponse, error) {

	var (
		resp   = &models.GetListPhoneResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_id,
			phone,
			descriprion,
			is_fax
		FROM phones
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			user_id     sql.NullString
			phone       sql.NullString
			descriprion sql.NullString
			is_fax      bool
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&phone,
			&descriprion,
			&is_fax,
		)

		if err != nil {
			return nil, err
		}

		resp.Phones = append(resp.Phones, &models.Phone{
			Id:          id.String,
			UserId:      user_id.String,
			Phone:       phone.String,
			Description: descriprion.String,
			IsFax:       is_fax,
		})
	}

	return resp, nil
}

func (r *phoneRepo) Update(ctx context.Context, req *models.UpdatePhone) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			phones
		SET
			phone = $1,
			descriprion = $2,
			is_fax = $3
		WHERE id = $4
	`

	result, err := r.db.Exec(ctx, query,
		req.Phone,
		req.Description,
		req.IsFax,
		req.Id,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *phoneRepo) Delete(ctx context.Context, req *models.PhonePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM phones WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
