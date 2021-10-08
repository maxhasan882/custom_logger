package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Log struct {
	Context *gin.Context
}

func (l Log) Println(message string, err error) {
	logString := fmt.Sprintf("[%v] ", l.Context.Request.Header.Get("Authorization"))
	if message != "" {
		logString += message
	}
	if err != nil {
		logString += err.Error()
	}
	log.Println(logString)
}
