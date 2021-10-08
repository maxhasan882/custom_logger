package logger

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Log struct {
	Context *gin.Context
}

func (l Log) Println(userID string, requestID string, err error) {
	log.Printf("[%v:%v] %v", userID, requestID, err)
	log.Println(l.Context.Request.Header.Get("Authorization"))
}
