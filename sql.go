package main

import (
    "fmt"
    "os"
    //"net"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type MyDB struct {
    Db		*sql.DB
    DbFile      string
    DbExists    bool
    DbConnected bool
}

func NewMyDb() *MyDB {
    return &MyDB{DbFile: "simple.sqlite"}

}

func (this *MyDB)Open() bool {
    if this.DbConnected == false {
	var err error
        this.Db, err = sql.Open("sqlite3", this.DbFile)
        if err != nil {
            fmt.Printf("Error: failed to open DB %s ",this.DbFile)
            fmt.Println(" with error ", err.Error())
            return false
        }
        this.DbConnected = true
    }
    if this.DbExists == false {
        if this.Create() == false {
	    fmt.Printf("Error: failed to Setup opened DB %s ",this.DbFile)
            return false
        }
    }
    return true
}
func (this *MyDB)Create() bool {
	stats, err := os.Stat(this.DbFile)
	if err != nil {
		_, err = this.Db.Exec("CREATE TABLE data (id TEXT not null primary key, content TEXT);")
		if err != nil {
			fmt.Println("Error: failed to create db", err.Error())
			return false
		}
		this.DbExists=true
	    
	} else {
		fmt.Println("File exists, stats: ",stats)
		this.DbExists=true
	}
	return this.DbExists
}
func (this *MyDB)Close() *MyDB {
	this.Db.Close()
	return nil
}
