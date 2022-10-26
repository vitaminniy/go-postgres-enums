package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

var (
	db_user     = os.Getenv("DB_USER")
	db_password = os.Getenv("DB_PASSWORD")
	db_name     = os.Getenv("DB_NAME")
	db_host     = os.Getenv("DB_HOST")
)

func main() {
	log.SetFlags(0)

	runsuite("pgx")
	runsuite("postgres")
}

func runsuite(driver string) {
	log.SetPrefix(driver + ":")

	db, err := sql.Open(driver, connstr())
	if err != nil {
		log.Fatalf("could not open db with: %v", driver, err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("could not establish db conn: %v", err)
	}

	runintssuite(db)
	runstringssuite(db)
}

func connstr() string {
	conn := "postgresql://"
	conn = conn + db_user
	conn = conn + ":"
	conn = conn + db_password
	conn = conn + "@"
	conn = conn + db_host
	conn = conn + ":5432/"
	conn = conn + db_name
	conn = conn + "?sslmode=disable"

	return conn
}

func runintssuite(db *sql.DB) {
	if err := truncate(db); err != nil {
		log.Fatalf("could not truncate db: %v", err)
	}

	if err := insertIntColors(db); err != nil {
		log.Fatalf("could not insert int colors: %v", err)
	}

	lightints, err := readIntColors(db)
	if err != nil {
		log.Fatalf("could not read int colors: %v", err)
	}

	for _, li := range lightints {
		log.Printf("got light int: id=%d; color=%q", li.id, li.color)
	}
}

func runstringssuite(db *sql.DB) {
	if err := truncate(db); err != nil {
		log.Fatalf("could not truncate db: %v", err)
	}

	if err := insertStringColors(db); err != nil {
		log.Fatalf("could not insert string colors: %w", err)
	}

	lightstrings, err := readStringColors(db)
	if err != nil {
		log.Fatalf("could not read string colors: %v", err)
	}

	for _, ls := range lightstrings {
		log.Printf("got light string: id=%d; color=%q", ls.id, ls.color)
	}
}

func truncate(db *sql.DB) error {
	_, err := db.Exec("TRUNCATE lights")
	return err
}
