package postgres

import "github.com/jackc/pgtype/pgxtype"

type DBSettings interface {
	toDSN() string
}

type Store interface {
	getConnectionPool() pgxtype.Querier
}

type DBConnection interface {
	New(settings DBSettings) (*Store, error)
}
