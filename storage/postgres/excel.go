package postgres

import (
	"app/api/models"
	"context"
	"database/sql"
)

func (u *userRepo) GetAllStudentsForExcel(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp  = &models.UserGetListResponse{}
		query string
		where = " "
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			u.id,
			u.role_id,
			u.first_name,
			u.last_name,
			u.email,
			u.phone_number,
			u.password,
			u.created_at,
			u.updated_at,

			r.type
		FROM "users" AS u
		JOIN "roles" AS r ON u.role_id = r.id
	`

	if len(req.Filter) > 0 {
		where = " WHERE r.type = '" + req.Filter + "'"
	}

	query += where

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			roleId      sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			email       sql.NullString
			phoneNumber sql.NullString
			password    sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
			userType    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&roleId,
			&firstName,
			&lastName,
			&email,
			&phoneNumber,
			&password,
			&createdAt,
			&updatedAt,
			&userType,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:          id.String,
			RoleId:      roleId.String,
			FirstName:   firstName.String,
			LastName:    lastName.String,
			Email:       email.String,
			PhoneNumber: phoneNumber.String,
			Password:    "",
			UserType:    userType.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}
