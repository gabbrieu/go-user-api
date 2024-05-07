package exception

import "log"

func FatalLogging(err error, msg string) {
	if err != nil {
		log.Fatalf(msg)
	}
}
