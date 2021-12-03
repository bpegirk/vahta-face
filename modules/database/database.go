package database

import (
	"database/sql"
	"example.org/modules/config"
	"fmt"
)

var db *sql.DB

func getDb() *sql.DB {
	if nil == db {
		connStr := fmt.Sprintf("%s:%s@%s:%v/%s",
			configuration.Db.User,
			configuration.Db.Password,
			configuration.Db.Host,
			configuration.Db.Port,
			configuration.Db.Database,
		)
		db, _ = sql.Open("firebirdsql", connStr)
	}
	return db
}

func hardwareName(id int) string {
	var name string
	_db := getDb()
	err := _db.QueryRow("SELECT H_NAME FROM HARDWARE WHERE H_ID=?", id).Scan(&name)
	if err != nil {
		return ""
	}

	return name
}

func peopleName(id int) (string, []byte) {
	var name string
	var photo []byte
	_db := getDb()
	err := _db.QueryRow("SELECT P_FIO,P_FOTO FROM PEOPLE WHERE P_ID=?", id).Scan(&name, &photo)
	if err != nil {
		return "", nil
	}

	return name, photo
}
