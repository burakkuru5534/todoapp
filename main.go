package main

import (
	"database/sql"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const  (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "tayitkan"
	dbname = "todo_db"
)


func main () {

	//router
	r := mux.NewRouter()
	//api endpoints
	r.Handle("/todo", toDoCreate())
	r.Handle("/todo/update", toDoUpdate())
	r.Handle("/todo/delete", toDoDelete())

	//define options
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	//start server
	log.Fatal(http.ListenAndServe(":8081", corsWrapper.Handler(r)))

}

func toDoCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// connection string
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Errorf("%v",err)
		}
		// close database
		defer db.Close()

		// check db
		err = db.Ping()

		fmt.Println("Connected!")

		eventName := r.FormValue("eventName")
		sq := fmt.Sprintf(`insert into event (event_name) values('%s')`,eventName)

		_,err = db.Exec(sq)
		if err != nil {
			fmt.Errorf("%v",err)
			panic(err)
		}
	})
}

func toDoUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// connection string
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Errorf("%v",err)
			panic(err)
		}
		// close database
		defer db.Close()

		// check db
		err = db.Ping()

		fmt.Println("Connected!")

		eventName := r.FormValue("eventName")
		eventID := r.FormValue("eventID")

		eventIDint, err := strconv.ParseInt(eventID, 10, 64)
		if err != nil {
			panic(err)
		}

		sq := fmt.Sprintf(`update event set event_name = '%s' where id = %d`,eventName, eventIDint)

		_,err = db.Exec(sq)
		if err != nil {
			fmt.Errorf("%v",err)
			panic(err)
		}
	})
}

func toDoDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// connection string
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Errorf("%v",err)
		}
		// close database
		defer db.Close()

		// check db
		err = db.Ping()

		fmt.Println("Connected!")

		eventID := r.FormValue("eventID")

		eventIDint, err := strconv.ParseInt(eventID, 10, 64)
		if err != nil {
			panic(err)
		}

		sq := fmt.Sprintf(`delete from event where id = %d`, eventIDint)

		_,err = db.Exec(sq)
		if err != nil {
			fmt.Errorf("%v",err)
			panic(err)
		}
	})
}