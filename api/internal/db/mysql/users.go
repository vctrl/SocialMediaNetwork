package mysql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/vctrl/social-media-network/api/internal/config"
)

type UsersMySQL struct {
	db *sql.DB
}

func NewUsersMySQL(db *sql.DB) Users {
	return &UsersMySQL{
		db: db,
	}
}

func FromConfig(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.MySQL.Conn)
	if err != nil {
		return nil, errors.Wrap(err, "open sql connection")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping db")
	}

	return db, nil
}

func (r *UsersMySQL) Add(ctx context.Context, user *User) error {
	query := "INSERT INTO users (`id`, `login`, `password`, `created_at`) VALUES (?, ?, ?, ?)"
	res, err := r.db.Exec(query, user.ID, user.Login, user.Password, user.CreatedAt)
	if err != nil {
		return errors.Wrap(err, "insert user")
	}

	_, err = res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "get last insert id")
	}

	return nil
}

func (r *UsersMySQL) GetByLogin(ctx context.Context, login string) (*User, error) {
	query := "SELECT * FROM users WHERE login = ?"

	res := r.db.QueryRow(query, login)

	u := &User{}
	err := res.Scan(&u.ID, &u.Login, &u.Password, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "scan user row")
	}

	return u, nil
}
