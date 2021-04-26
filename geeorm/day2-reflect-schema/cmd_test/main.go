package main

import (
	"fmt"
	"geeorm"
	"geeorm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := geeorm.NewEngine("sqlite3", "gee.db")
	if err != nil {
		log.Error(err)
	}
	defer engine.Close()
	session := engine.NewSession()
	_, _ = session.Raw("DROP TABLE IF EXISTS USER;").Exec()
	_, _ = session.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = session.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := session.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
