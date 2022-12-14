# go-socket

you can send text with `nc -U echo.sock`

## install

`go install github.com/test3-damianfurrer/go-socket@latest`

### sql server example
`go install github.com/test3-damianfurrer/go-socket@dbServer`
 
  - creates(if inexistent) a local sqlite3 database
  - data persistent in the db file
  - interprets 'read <id>' or 'write <id> <text>' commands
  - allows multiple instances to read from or write to the db  
