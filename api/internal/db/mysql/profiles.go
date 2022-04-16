package mysql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type ProfilesMySQL struct {
	db *sql.DB
}

func NewProfilesMySQL(db *sql.DB) Profiles {
	return &ProfilesMySQL{db: db}
}

func (r *ProfilesMySQL) GetByUserID(ctx context.Context, id string) (*Profile, error) {
	query := "SELECT * FROM profiles WHERE user_id  = ?"

	res := r.db.QueryRow(query, id)
	p := &Profile{}
	err := res.Scan(&p.ID, &p.UserID, &p.Name, &p.Surname, &p.Sex, &p.Interests, &p.City)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "scan profile row")
	}

	return p, nil
}
