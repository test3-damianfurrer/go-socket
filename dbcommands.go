package main

import(
    "fmt"
    "net"
    "strings"
)

func prcdbcommand(databuf []byte,conn net.Conn,db *MyDB) {
    command := fmt.Sprintf("%s",databuf)
    carr := strings.Split(command," ")
    for i := 0; i < len(carr); i++ {
        fmt.Printf("Index %d: %s",i,carr[i])
    }
    
    ///tmp
    if conn == nil {
        fmt.Println("conn nil")
    }
    if db == nil {
        fmt.Println("db nil")
    }
}
