package main

import(
    "fmt"
    "net"
    "strings"
)

func prcdbcommand(databuf []byte,conn net.Conn,mydb *MyDB) {
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
            readDB(actnid, conn, mydb.Db)
        case "WRITE":
            writeDB(actnid, text, conn, mydb.Db)
    }
}

func readDB(id string, conn net.Conn,db *sql.DB) {
    result, err := db.Query("select content from data where id like ?",id + "%")
    if err != nil {
        fmt.Println("Query Error: ", err.Error())
    }
    var id string
    var content string
    for rows.Next() {
        err = rows.Scan(&id, &content)
        if err != nil {
            fmt.Println("Error reading rows: ", err.Error())
        } else {
            sendstring := fmt.Sprintf("id : %s; content: %s", id, content)
            fmt.Printf("DEBUG: %s", sendstring)
        }
        conn.Write([]byte{'x','\n'})
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
