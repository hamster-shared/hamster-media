package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

// Fail failed result
func Fail(message string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Result{
		Code:    ERROR,
		Message: message,
		Data:    nil,
	})
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code:    SUCCESS,
		Message: "success",
		Data:    data,
	})
}
