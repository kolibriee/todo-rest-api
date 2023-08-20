package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tryrest "github.com/kolibri7557/try-rest-api"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input tryrest.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ListId, err := h.services.TodoList.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": ListId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
}

func (h *Handler) getListById(c *gin.Context) {
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
