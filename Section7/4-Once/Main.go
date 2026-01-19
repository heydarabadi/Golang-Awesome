package main

import "sync"

var once sync.Once

type Config struct {
	ConnectionString string
}

var configTest Config

func main() {
	for i := 0; i < 100; i++ {
		GetConfig()
	}
}

func GetConfig() Config {
	once.Do(func() {
		configTest = Config{
			ConnectionString: "Pouya",
		}
		println("GetConfig -- ", &configTest)
	})

	println("GetConfig Old -- ", &configTest)
	return configTest
}
