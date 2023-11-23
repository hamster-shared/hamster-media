package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) FaucetList(gin *gin.Context) {
	data, err := h.faucetService.FaucetList()
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
