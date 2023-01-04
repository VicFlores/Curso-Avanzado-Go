package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

func (DataBase) CreateSingleConnection() {
	fmt.Println("Creating singleton for database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation done")
}

var db *DataBase

var lock sync.Mutex

func GetDataBaseInstance() *DataBase {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating DB connection")
		db = &DataBase{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB already created")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			GetDataBaseInstance()
		}()
	}

	wg.Wait()
}
