package mdb

import (
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

type EmailEntry struct {
	Id          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool
}

// Create new database. This function will abort the program if
// there is no database and it failed to create one.
func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE emails (
			id           INTEGER PRIMARY KEY,
			email        TEXT UNIQUE,
			confirmed_at TIMESTAMP,
			opt_out      INTEGER
		);
	`)
	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			// abort on all sqlite error codes except code 1: "table already exists"
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else {
			// abort on all other error types
			log.Fatal(err)
		}
	}
}

// Build an EmailEntry from a database row
func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var id int64
	var email string
	var confirmedAt *time.Time
	var optOut bool

	err := row.Scan(&id, &email, &confirmedAt, &optOut)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &EmailEntry{Id: id, Email: email, ConfirmedAt: confirmedAt, OptOut: optOut}, nil

}

func GetEmail(db *sql.DB, email string) (*EmailEntry, error) {
	rows, err := db.Query(`SELECT id, email, confirmed_at, opt_out FROM emails WHERE email = ?`, email)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return emailEntryFromRow(rows)
	}
	return nil, nil
}

func UpdateEmail(db *sql.DB, entry EmailEntry) error {
	_, err := db.Exec(`INSERT INTO
		emails(email, confirmed_at, opt_out)
		VALUES(?, ?, ?)
		ON CONFLICT(email) DO UPDATE SET
		  confirmed_at=?,
		  opt_out=?`, entry.Email, entry.ConfirmedAt, entry.OptOut, entry.ConfirmedAt, entry.OptOut)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

type GetAllEmailsQueryParams struct {
	Page  int
	Count int
}

func GetAllEmails(db *sql.DB, params GetAllEmailsQueryParams) ([]EmailEntry, error) {
	var empty []EmailEntry
	rows, err := db.Query(`SELECT
		id, email, confirmed_at, opt_out
		FROM emails
		WHERE opt_out = false
		ORDER BY id ASC
		LIMIT ? OFFSET ?`, params.Count, (params.Page-1)*params.Count)
	if err != nil {
		log.Print(err)
		return empty, err
	}

	defer rows.Close()

	emails := make([]EmailEntry, 0, params.Count)

	for rows.Next() {
		email, err := emailEntryFromRow(rows)
		if err != nil {
			return nil, err
		}
		emails = append(emails, *email)
	}

	return emails, nil
}
