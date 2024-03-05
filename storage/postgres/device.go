package postgres

import (
	"app/api/models"
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type deviceRepo struct {
	db *pgxpool.Pool
}

func NewDeviceRepo(db *pgxpool.Pool) *deviceRepo {
	return &deviceRepo{
		db: db,
	}
}

func (u *deviceRepo) Create(ctx context.Context, req *models.DeviceCreate) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "devices"(id, user_id, device_name, browser, browser_version, ip)
		VALUES ($1, $2, $3, $4, $5,$6)
	`

	_, err := u.db.Exec(ctx, query,
		id,
		req.UserId,
		req.DeviceName,
		req.Browser,
		req.BrowserVersion,
		req.IP,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *deviceRepo) GetList(ctx context.Context, req *models.DeviceGetListRequest) (*models.DeviceGetListResponse, error) {

	var (
		resp   = &models.DeviceGetListResponse{}
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_id,
			device_name,
			browser,
			browser_version,
			ip,
			created_at
		FROM "devices" 
		WHERE user_id = $1
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit

	rows, err := u.db.Query(ctx, query, req.UserId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id             sql.NullString
			userId         sql.NullString
			deviceName     sql.NullString
			browser        sql.NullString
			browserVersion sql.NullString
			ip             sql.NullString
			createdAt      sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&userId,
			&deviceName,
			&browser,
			&browserVersion,
			&ip,
			&createdAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Devices = append(resp.Devices, &models.Device{
			Id:             id.String,
			UserId:         userId.String,
			DeviceName:     deviceName.String,
			Browser:        browser.String,
			BrowserVersion: browserVersion.String,
			IP:             ip.String,
			CreatedAt:      createdAt.String,
		})
	}
	return resp, nil
}

func (u *deviceRepo) Delete(ctx context.Context, req *models.DevicePrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE FROM "devices" WHERE id = $1`, req.Id)
	if err != nil {
		return err
	}

	return nil
}
