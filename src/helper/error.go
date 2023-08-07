package helper

import "log"

func CheckAndLogError(errMessage string, err error) {
	if err != nil {
		log.Fatalf(errMessage, err)
	}
}

func CheckMatchAndLogError(errMessage string, err string, match bool) {
	if match == false {
		log.Fatalf(errMessage, err)
	}
}
