package utils

import (
	"log"
)

func HasErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HasError(err error) {
	if err != nil {
		log.Println(err)
	}
}
