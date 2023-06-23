package handler

import (
	"net/http"

	"github.com/Cadeusept/notes-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input notes.NoteList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.NoteList.Create(user_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
