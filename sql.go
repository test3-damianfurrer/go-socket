package main

import (
    "fmt"
    "os"
	"net"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type MyDB struct {
    DbFile      string
    DbExists    bool
    DbConnected bool
}

func NewMyDb() *MyDB {
    return &MyDB{DbFile: "simple.sqlite"}

}

func (this *MyDB)Open() bool {
    if this.DbExists == false {
        if this.Create() == false {
            return false
        }
    }
    if this.DbConnected == false {
        db, err := sql.Open("sqlite3", this.DbFile)
        if err != nil {
            fmt.Printf("Error: failed to open DB %s ",this.DbFile)
            fmt.Println(" with error ", err.Error())
            return false
        }
        this.DbConnected = true
    }
    return true
}
func (this *MyDB)Create() bool {
    return true
}
