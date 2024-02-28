package postgres

import (
	"app/api/models"
	"context"
	"database/sql"

	uuid "github.com/google/uuid"
)

func (u *courseRepo) UploadPhotos(ctx context.Context, req *models.VideoLessons) (string, error) {

	var (
		id string = uuid.New().String()
	)

	insertQuery := `
            INSERT INTO "photos" (id, name, data) VALUES ($1, $2, $3)
        `
	_, err := u.db.Exec(ctx, insertQuery, id, req.FileName, req.PhotoData)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *courseRepo) GetPhotos(ctx context.Context, req *models.VideoLessons) (*models.VideoLessons, error) {

	var (
		id   sql.NullString
		name sql.NullString
		data []byte
	)

	insertQuery := `
            SELECT 
				id,
				name,
				data
			FROM "photos"
			WHERE id = $1
        `
	err := u.db.QueryRow(ctx, insertQuery, req.Id).Scan(
		&id,
		&name,
		&data,
	)
	if err != nil {
		return nil, err
	}

	return &models.VideoLessons{
		Id:        id.String,
		FileName:  name.String,
		PhotoData: data,
	}, nil
}

