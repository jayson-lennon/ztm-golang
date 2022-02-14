package main

import (
	"database/sql"
	"log"
	"mailinglist/api"
	"mailinglist/mdb"
	"net/http"

	"github.com/alexflint/go-arg"
)

// Flags that may be passed to the program
var args struct {
	DbPath string `arg:"env:MAILINGLIST_DB"`
	Bind   string `arg:"env:MAILINGLIST_BIND"`
}

func main() {
	arg.MustParse(&args)

	// Defaults if none provided
	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.Bind == "" {
		args.Bind = ":8080"
	}

	log.Printf("using database '%v'\n", args.DbPath)
	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Try to create database if it doesn't exist.
	mdb.TryCreate(db)

	http.Handle("/email/get", api.GetEmail(db))
	http.Handle("/email/get_all", api.GetAllEmails(db))
	http.Handle("/email/update", api.UpdateEmail(db))
	http.Handle("/email/delete", api.DeleteEmail(db))
	log.Printf("listening on %v\n", args.Bind)
	log.Print(http.ListenAndServe(args.Bind, nil))
}
