package main

import (
	"log"
	"os"
)

func main() {
	sum(2, 3)
}

var (
	errorLogger  *log.Logger
	warningLoger *log.Logger
	infoLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|
		os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Filed To Open Log File", err)
	}

	flags := log.Llongfile | log.LstdFlags | log.Lshortfile | log.Ltime | log.Lmicroseconds

	log.SetFlags(flags)
	log.SetOutput(file)

	warningLoger = log.New(file, "Warning! ", flags)
	infoLogger = log.New(file, "Info! ", flags)
	errorLogger = log.New(file, "Error! ", flags)

}

func sum(a, b int) {
	errorLogger.Println("Sum Started")
	println(a + b)
	infoLogger.Println("Sum Ended")
}
