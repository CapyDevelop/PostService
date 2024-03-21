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
		"host=%s  user=%s ",
		host, user)
	if password != "" {
		dsn += fmt.Sprintf("password=%s ", password)
	}
	dsn += fmt.Sprintf("dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", database, port)

	fmt.Println(dsn)

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
