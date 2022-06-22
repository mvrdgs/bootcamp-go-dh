package mysqlStore

import (
	"database/sql"
	"log"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:@tcp(localhost:3306)/storage"
	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = storageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println()
}
