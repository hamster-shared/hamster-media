package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *HandlerServer) GetNavbarList(gin *gin.Context) {
	vos, err := h.navbarService.GetNavbarList()
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(vos, gin)
}

func (h *HandlerServer) GetNavbarContent(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	vos, err := h.navbarService.GetNavbarContent(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(vos, gin)
}
