package config

import (
	"fmt"
	"sync"
)

var once sync.Once

type config struct {
}

func (cf config) Info() {
	fmt.Println("DBUsername:postgres, DBPort: 5432, DBName: postgres")
}

var singleInstance *config

func GetInstanceOfDB() *config {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &config{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
