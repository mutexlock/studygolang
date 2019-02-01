package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// DB mysql db pool
	DB *sqlx.DB
)

type dbConfigItem struct {
	Host    string
	Port    int
	User    string
	Passwd  string
	Db      string
	Charset string
}

func main() {
	mapMember := make(map[string]interface{})
	qukanConf := new(dbConfigItem)
	qukanConf.User = "root"
	qukanConf.Port = 3306
	qukanConf.Host = "127.0.0.1"
	qukanConf.Db = "test"
	qukanConf.Passwd = ""
	qukanConf.Charset = "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", qukanConf.User, qukanConf.Passwd, qukanConf.Host, qukanConf.Port, qukanConf.Db, qukanConf.Charset)
	DB = sqlx.MustConnect("mysql", dsn)
	DB.SetMaxOpenConns(256)
	DB.SetMaxIdleConns(32)
	DB.SetConnMaxLifetime(600 * time.Second)

	if DB != nil {
		rows, err := DB.Queryx("SELECT info FROM member1 where id =1 limit 1")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()
		rows.Next()
		rows.MapScan(mapMember)
		//rows.StructScan(&mapMember)
		fmt.Println(mapMember["info"])
	}

	var names []string
	err := DB.Select(&names, "SELECT info FROM member1 where id =3 limit 1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)

	var name string

	err = DB.Get(&name, "SELECT info FROM member1 where id =? limit 1", 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name)

}
