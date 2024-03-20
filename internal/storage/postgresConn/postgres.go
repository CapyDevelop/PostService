package postgresConn

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func New(user, password, database, host string, port int) (*sql.DB, error) {
	op := "storage.postgres.New"
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, database, port)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
