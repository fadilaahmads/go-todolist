package utils

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
