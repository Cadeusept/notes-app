package handler

import (
	"net/http"
	"strconv"

	"github.com/Cadeusept/notes-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list_id, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id")
		return
	}

	var input notes.NoteItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.NoteItem.Create(user_id, list_id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []notes.NoteItem `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list_id, err := strconv.Atoi(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id")
		return
	}

	items, err := h.services.NoteItem.GetAll(user_id, list_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(*gin.Context) {

}

func (h *Handler) updateItem(*gin.Context) {

}

func (h *Handler) deleteItem(*gin.Context) {

}
