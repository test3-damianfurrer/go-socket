package main

import (
    "fmt"
    "os"
	"net"
//	"io"
)

func echoServer(c net.Conn) {
    fmt.Printf("Client connected [%s]\n", c.RemoteAddr().Network())
    fmt.Println("addr",c.RemoteAddr())
    //io.Copy(c, c)
	for{
		databuf := make([]byte,0)
		tmpbuf := make([]byte, 1)
		
		for {
			_, err := c.Read(tmpbuf)
			if err != nil {
				if err.Error() != "EOF"{
					fmt.Println("READ ERR",err.Error())
				} else {
					fmt.Println("Client Connection Closed",err.Error())
				}
				c.Close()
				return
					
			}
			//fmt.Println("byte",tmpbuf[0])
			if tmpbuf[0] == '\n' {
				databuf = append(databuf,'\n')
				break
			}
			if tmpbuf[0] == 0 {
				databuf = append(databuf,'\n')
				break
			}
			if tmpbuf[0] == 10 {
				databuf = append(databuf,'\n')
				break
			}
			databuf = append(databuf,tmpbuf[0])
		}
		fmt.Printf("Received: %s",databuf)
		c.Write([]byte{'Y','o','u',' ','s','e','n','t',':',' '})
		c.Write(databuf)
	}
    c.Close()
    fmt.Println("Connection Closed")
}

func dbServer(c net.Conn, db *MyDB) {
    fmt.Printf("Client connected [%s]\n", c.RemoteAddr().Network())
    fmt.Println("addr",c.RemoteAddr())
    //io.Copy(c, c)
	for{
		databuf := make([]byte,0)
		tmpbuf := make([]byte, 1)
		
		for {
			_, err := c.Read(tmpbuf)
			if err != nil {
				if err.Error() != "EOF"{
					fmt.Println("READ ERR",err.Error())
				} else {
					fmt.Println("Client Connection Closed",err.Error())
				}
				c.Close()
				return
					
			}
			//fmt.Println("byte",tmpbuf[0])
			if tmpbuf[0] == '\n' {
				//databuf = append(databuf,'\n')
				break
			}
			if tmpbuf[0] == 0 {
				//databuf = append(databuf,'\n')
				break
			}
			if tmpbuf[0] == 10 {
				//databuf = append(databuf,'\n')
				break
			}
			databuf = append(databuf,tmpbuf[0])
		}
		//fmt.Printf("Received: %s",databuf)
		//c.Write([]byte{'Y','o','u',' ','s','e','n','t',':',' '})
		//c.Write(databuf)
		//c net.Conn, db *MyDB
		prcdbcommand(databuf,c,db)
		
	}
    c.Close()
    fmt.Println("Connection Closed")
}

func main() {
    mydir, err := os.Getwd()
    if err != nil {
        fmt.Println("Can't get Current Directory",err.Error())
		return
    }
	SockAddr:=mydir + "/echo.sock"

    if err := os.RemoveAll(SockAddr); err != nil {
        panic(err)
    }

    mydb:=NewMyDb()
	mydb.DbFile = mydir+"/test.sqlite"
    if mydb.Open() {
		fmt.Println("DB opened")
	} else {
		fmt.Println("DB NOT opened")
		return
	}
	defer mydb.Close()
	//defer mydb = mydb.Close()
	
	
    l, err := net.Listen("unix", SockAddr)
    if err != nil {
        fmt.Println("listen error:",err.Error())
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("accept error:", err.Error())
        }

        //go echoServer(conn)
		go dbServer(conn, mydb)
    }
}
