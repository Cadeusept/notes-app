package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(*gin.Context) {

}

func (h *Handler) getListById(*gin.Context) {

}

func (h *Handler) updateList(*gin.Context) {

}

func (h *Handler) deleteList(*gin.Context) {

}
