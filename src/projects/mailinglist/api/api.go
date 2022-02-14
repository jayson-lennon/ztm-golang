package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"mailinglist/mdb"
	"net/http"
)

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// Populate the `target` data stucture from the JSON payload in `body`
func fromJson(body io.Reader, target interface{}) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	json.Unmarshal(buf.Bytes(), &target)
}

// Writes some data into a ResponseWriter in JSON format.
// Automatically performs the JSON conversion via the getData function.
// Any errors returned from the getData function will be considered a server error
// and status 500 will be returned along with the error message
func returnJson(w http.ResponseWriter, getData func() (interface{}, error)) {
	// set the correct content-type for all JSON responses
	setJsonHeader(w)

	data, serverErr := getData()

	// handle any errors with the getData function
	if serverErr != nil {
		// when a server error occurs, we set the status code to 500
		w.WriteHeader(500)
		serverErrJson, err := json.Marshal(&serverErr)
		if err != nil {
			// Not much we can do at this point, just log and abort with a 500.
			log.Print(err)
			return
		}
		// write the JSON formatted error into the stream
		w.Write(serverErrJson)
		return
	}

	// data is good at this point, try to serialize it
	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Write(dataJson)
}

// Handler for /email/get
func GetEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		returnJson(w, func() (interface{}, error) {
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

// Handler for /email/update
func UpdateEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		mdb.UpdateEmail(db, entry)

		returnJson(w, func() (interface{}, error) {
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

// Handler for /email/delete
func DeleteEmail(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			return
		}
		entry := mdb.EmailEntry{}
		fromJson(req.Body, &entry)

		// Important! Always "opt out" the email (instead of delete)
		// to prevent it from being used in non-transactional messaging.
		entry.OptOut = true
		mdb.UpdateEmail(db, entry)

		returnJson(w, func() (interface{}, error) {
			return mdb.GetEmail(db, entry.Email)
		})
	})
}

// Handler for /email/get_all
func GetAllEmails(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// only response to GET requests
		if req.Method != "GET" {
			return
		}

		queryOptions := mdb.GetAllEmailsQueryParams{}
		fromJson(req.Body, &queryOptions)

		// make sure we have the required fields
		if queryOptions.Count <= 0 || queryOptions.Page <= 0 {
			// If we don't have both fields, then return informational message
			// indicating proper usage of the API.
			returnJson(w, func() (interface{}, error) {
				errorMessage := struct {
					Err     string
					Example string
				}{
					Err:     "Page and Count fields are required and must be > 0",
					Example: `{"Page":1,"Count":10}`,
				}
				w.WriteHeader(400) // 400 == client error
				return errorMessage, nil
			})
			return
		}

		// everything check outs, run the query
		returnJson(w, func() (interface{}, error) {
			return mdb.GetAllEmails(db, queryOptions)
		})
	})
}
