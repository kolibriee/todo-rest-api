package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tryrest "github.com/kolibri7557/try-rest-api"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input tryrest.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.CreateItem(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id parameter")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id parameter")
		return
	}
	item, err := h.services.TodoItem.GetItemById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input tryrest.TodoItemUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.TodoItem.UpdateItem(userId, itemId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.DeleteItem(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
