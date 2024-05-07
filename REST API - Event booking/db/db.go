package db

// import (
// 	"database/sql"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var DB *sql.DB

// func InitDb() {
// 	var err error
// 	DB, err = sql.Open("sqlite3", "api.db")

// 	if err != nil {
// 		panic("Could not connect to database")
// 	}

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	createTables()
// }

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Open("mysql", "root:Harsh@Mysql@2001@tcp(127.0.0.1:3306)/GO_REST_Event_Booking?parseTime=true")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL UNIQUE,
        password TEXT NOT NULL
    )
    `

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}
