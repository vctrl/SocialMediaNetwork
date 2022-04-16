package mysql

import (
	"database/sql"
)

type FriendsMySQL struct {
	db *sql.DB
}

func NewFriendsMySQL(db *sql.DB) Friends {
	return &FriendsMySQL{
		db: db,
	}
}

type FriendRequestsMySQL struct {
	db *sql.DB
}

func NewFriendRequestsMySQL(db *sql.DB) FriendRequests {
	return &FriendRequestsMySQL{
		db: db,
	}
}
