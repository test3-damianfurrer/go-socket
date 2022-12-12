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
				fmt.Println("READ ERR")
				//break
				return
			}
			fmt.Println("byte",tmpbuf[0])
			if tmpbuf[0] == '\n' {
				databuf = append(databuf,'\n')
				break
			}
			if tmpbuf[0] == 0 {
				databuf = append(databuf,'\n')
				break
			}
			databuf = append(databuf,tmpbuf[0])
		}
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

        go echoServer(conn)
    }
}
