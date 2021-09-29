package util

import "log"

func PanicIfError(err error, message string) {
	if err != nil {
		log.Panicf("Failed: %s [%s]", message, err)
	}
}
