package utils

import "log"

func Logger(l ...interface{}) {
	log.Println(l)
}
