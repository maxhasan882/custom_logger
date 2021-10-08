package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Log struct {
	Context *gin.Context
}

func (l Log) Println(message string, err error) {
	logString := fmt.Sprintf("[%v] ", splitToken(l.Context.Request.Header.Get("Authorization")))
	if message != "" {
		logString += message
	}
	if err != nil {
		logString += err.Error()
	}
	log.Println(logString)
}

func splitToken(token string) string {
	splitData := strings.Split(token, " ")
	if len(splitData) != 2 {
		return ""
	}
	return splitData[1]
}
