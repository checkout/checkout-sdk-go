package configuration

import (
	"log"
	"os"
)

type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

func DefaultLogger() *log.Logger {
	return log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags)
}
