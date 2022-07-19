package repository

import "database/sql"

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}

func (repo *SQLiteRepository) Mirate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		amout real not null,
		purchase_date integer not null,
		purchase_price integer not null);
	`

	_, err := repo.Conn.Exec(query)

	return err
}
