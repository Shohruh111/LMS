package postgres

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *phoneRepo) Create(ctx context.Context, req *models.CreatePhone) (bool, error) {

	var (
		query string
	)

	query = `
		INSERT INTO phones(user_id, phone, descriprion, is_fax)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		req.UserId,
		req.Phone,
		req.Description,
		req.IsFax,
	)

	if err != nil {
		return false, err
	}

	return true, nil
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
			description,
			is_fax
		FROM phones
		WHERE user_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.UserId).Scan(
		&user_id,
		&phone,
		&descriprion,
		&is_fax,
	)

	if err != nil {
		return nil, err
	}

	return &models.Phone{
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
		order  = "ORDER BY DESC"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			user_id,
			phone,
			description,
			is_fax
		FROM phones
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
			user_id     sql.NullString
			phone       sql.NullString
			descriprion sql.NullString
			is_fax      bool
		)

		err := rows.Scan(
			&resp.Count,
			&user_id,
			&phone,
			&descriprion,
			&is_fax,
		)

		if err != nil {
			return nil, err
		}

		resp.Phones = append(resp.Phones, &models.Phone{
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
			description = $2,
			is_fax = $3
		WHERE user_id = $4
	`

	result, err := r.db.Exec(ctx, query,
		req.Phone,
		req.Description,
		req.IsFax,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *phoneRepo) Delete(ctx context.Context, req *models.PhonePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM phones WHERE id = $1", req.UserId)
	if err != nil {
		return err
	}

	return nil
}
