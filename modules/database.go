package modules

import (
	"database/sql"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
)

var db *sql.DB

func getDb() *sql.DB {
	if nil == db {
		connStr := fmt.Sprintf("%s:%s@%s:%v/%s",
			Cfg.Db.User,
			Cfg.Db.Password,
			Cfg.Db.Host,
			Cfg.Db.Port,
			Cfg.Db.Database,
		)
		var err error
		db, err = sql.Open("firebirdsql", connStr)
		if err != nil {
			fmt.Println("Error connection to database: ", err)
			panic("Hold app")
		}
	}
	return db
}

func HardwareName(id int) string {
	var name string
	_db := getDb()
	row := _db.QueryRow("SELECT H_NAME FROM HARDWARE WHERE H_ID=?", id)
	if row.Err() != nil {
		fmt.Println("Error get hardware name: ", row.Err())
		return ""
	}
	err := row.Scan(&name)
	if err != nil {
		return ""
	}

	return name
}

func PeopleName(id int) (string, []byte) {
	var name string
	var photo []byte
	_db := getDb()
	err := _db.QueryRow("SELECT P_FIO,P_FOTO FROM PEOPLE WHERE P_ID=?", id).Scan(&name, &photo)
	if err != nil {
		return "", nil
	}

	return name, photo
}
