package helper

import "log"

func CheckAndLogError(errMessage string, err error) {
	if err != nil {
		log.Fatalf(errMessage, err)
	}
}
