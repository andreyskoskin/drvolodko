package users

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/andreyskoskin/drvolodko/domain"
)

type PgUsers struct {
	db *sql.DB
}

/*
func NewPgUsers(db *sql.DB) *PgUsers {
	return &PgUsers{db: db}
}
*/

func (pg *PgUsers) GetUser(id domain.UserID) (user domain.GetUserResponse, err error) {
	const query = `
		SELECT
			id,
			name
		FROM
			users
		WHERE id = $1
	`

	err = pg.db.QueryRow(query, string(id)).Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		err = domain.ErrNotFound
	}

	return user, err
}

func (pg *PgUsers) ListUsers() (_ domain.ListUsersResponse, err error) {
	var users []domain.ListUsersResponseItem

	const query = `
		SELECT
			id,
			name
		FROM
			users
	`

	var rows *sql.Rows
	rows, err = pg.db.Query(query)
	if err != nil {
		return domain.ListUsersResponse{}, err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var user domain.ListUsersResponseItem
		if err = rows.Scan(&user.ID, &user.Name); err != nil {
			return domain.ListUsersResponse{}, err
		}
	}

	return domain.ListUsersResponse{
		Items: users,
	}, rows.Err()
}

func (pg *PgUsers) AddUser(request domain.AddUserRequest) (domain.AddUserResponse, error) {
	var id, err = uuid.NewUUID()

	if err != nil {
		return domain.AddUserResponse{}, err
	}

	const query = `INSERT users(id, name) VALUES ($1, $2)`
	if _, err = pg.db.Exec(query, id, request.Name); err != nil {
		return domain.AddUserResponse{}, err
	}

	return domain.AddUserResponse{
		ID: domain.UserID(id.String()),
	}, nil
}

func (pg *PgUsers) KillUser(id domain.UserID) (err error) {
	_, err = pg.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}
