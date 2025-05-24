package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func insert(db *sql.DB,table string, columns []string, values []string) {
	fmt.Println("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")

	stmtIns, err := db.Prepare("insert into " + table + " (" + strings.Join(columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}
}