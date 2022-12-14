package main

import(
    "fmt"
    "net"
    "strings"
    "database/sql"
)

func prcdbcommand(databuf []byte,conn net.Conn,db *sql.DB) {
    maincmd := ""
    actnid := ""
    text := ""
    command := fmt.Sprintf("%s",databuf)
    carr := strings.Split(command," ")
    for i := 0; i < len(carr); i++ {
        fmt.Printf("Index %d: %s\n",i,carr[i])
        switch i {
            case 0:
                maincmd = strings.ToUpper(carr[0])
            case 1:
                actnid = strings.ToUpper(carr[1])
            default:
                if text == "" {
                    text = carr[i]
                } else {
                    text = text + " " + carr[i]
                }
        }
        
    }
    
    switch maincmd {
        case "READ":
            readDB(actnid, conn, db)
        case "WRITE":
            writeDB(actnid, text, conn, db)
    }
}

func readDB(id string, conn net.Conn,db *sql.DB) {
    rows, err := db.Query("select id, content from data where id like ?",id + "%")
    if err != nil {
        fmt.Println("Query Error: ", err.Error())
    }
    var tid string
    var tcontent string
    for rows.Next() {
        err = rows.Scan(&tid, &tcontent)
        if err != nil {
            fmt.Println("Error reading rows: ", err.Error())
        } else {
            sendstring := fmt.Sprintf("id : %s; content: %s\n", tid, tcontent)
            fmt.Printf("DEBUG: %s", sendstring)
            conn.Write(sendstring)
        }
        //conn.Write([]byte{'x','\n'})
    }
}

func writeDB(id string, text string, conn net.Conn,db *sql.DB) {
    fmt.Println("Debug str to insert: ",text)
    result, err := db.Exec("INSERT INTO data(id, content) VALUES (?, ?)",id, text)
    if err != nil {
        fmt.Println("Insert Error: ", err.Error())
    } else {
        fmt.Println("Inserted: ", result)
    }
    conn.Write([]byte{'y','\n'})
}
