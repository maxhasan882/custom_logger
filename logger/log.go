package logger

import "log"

func Println(userID string, requestID string, err error) {
	log.Printf("[%v:%v] %v", userID, requestID, err)
}
