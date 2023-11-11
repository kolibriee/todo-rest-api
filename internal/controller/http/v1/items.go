package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
)

// .
//
//	@Summary		Create item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	create todo item
//	@ID				create-item
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer					true	"list id"
//	@Param			input	body		domain.TodoItemCreate	true	"item info"
//	@Success		200		{integer}	integer					1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId}/items [post]
func (h *Handler) CreateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.TodoItemCreate
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

// .
//
//	@Summary		Get all items
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	get all items
//	@ID				get-all-items
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer	true	"list id"
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId}/items [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("listId"))
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

// .
//
//	@Summary		Get item by id
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	get item by id
//	@ID				get-item-by-id
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer	true	"list id"
//	@Param			itemId	path		integer	true	"item id"
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId}/items/{itemId} [get]
func (h *Handler) GetItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("itemId"))
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

// .
//
//	@Summary		Update item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	update item
//	@ID				update-item
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer					true	"list id"
//	@Param			itemId	path		integer					true	"item id"
//	@Param			input	body		domain.TodoItemUpdate	true	"item info"
//	@Success		200		{integer}	integer					1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId}/items/{itemId} [put]
func (h *Handler) UpdateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.TodoItemUpdate
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

// .
//
//	@Summary		Delete item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	delete item
//	@ID				delete-item
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer	true	"list id"
//	@Param			itemId	path		integer	true	"item id"
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId}/items/{itemId} [delete]
func (h *Handler) DeleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.TodoItem.DeleteItem(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
