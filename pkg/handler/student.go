package handler

import (
	"github.com/gin-gonic/gin"
	http "net/http"
)

func (h *Handler) createStudent(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getStudents(c *gin.Context) {

}

func (h *Handler) getStudent(c *gin.Context) {
}

func (h *Handler) updateStudent(c *gin.Context) {
}

func (h *Handler) deleteStudent(c *gin.Context) {

}
