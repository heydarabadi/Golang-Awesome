package main

import (
	"fmt"
	"sync"
)

type DbConnection struct {
	Host     string
	DbName   string
	User     string
	Password string
}

var connectionPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return &DbConnection{
			Host:     "localhost",
			DbName:   "Db",
			User:     "root",
			Password: "root",
		}
	},
}

func main() {
	connection := connectionPool.Get().(*DbConnection)
	fmt.Printf("%v \n", connection)

	connectionPool.Put(connection)

	connection = connectionPool.Get().(*DbConnection)
	fmt.Printf("%v \n", connection)
}
