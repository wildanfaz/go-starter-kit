package repositories

import "database/sql"

type ImplementUsersRepository struct {
	dbMySql *sql.DB
}

type UsersReporitory interface {
}

func NewUsersRepository(dbMySql *sql.DB) UsersReporitory {
	return &ImplementUsersRepository{
		dbMySql: dbMySql,
	}
}
