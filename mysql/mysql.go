package main

import (
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// Users ...
type Users struct {
	ID       int
	Name     string
	Password string
	Tel      string
}

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	
	for rows.Next() {
		// user := new(Users)  // 不得使用 new
		user := Users{}

		// err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Tel)

		// 如果列太多，一个一个枚举是不合适的，因此使用 reflect 遍历
		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface() // 取每个字段的指针地址
		}

		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
