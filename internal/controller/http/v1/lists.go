package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
)

// .
//
//	@Summary		Create list
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	create todo list
//	@ID				create-list
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.TodoListCreate	true	"list info"
//	@Success		200		{integer}	integer					1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists [post]
func (h *Handler) CreateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.TodoListCreate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listId, err := h.services.TodoList.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listId,
	})
}

// .
//
//	@Summary		Get all lists
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	get all lists
//	@ID				get-all-lists
//	@Accept			json
//	@Produce		json
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists [get]
func (h *Handler) GetAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lists)
}

// .
//
//	@Summary		Get list
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	get list
//	@ID				get-list
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		int		true	"list id"
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId} [get]
func (h *Handler) GetListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// .
//
//	@Summary		Update list
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	update list
//	@ID				update-list
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer					true	"list id"
//	@Param			input	body		domain.TodoListUpdate	true	"list info"
//	@Success		200		{integer}	integer					1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId} [put]
func (h *Handler) UpdateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.TodoListUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.TodoList.UpdateList(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// .
//
//	@Summary		Delete list
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	delete list
//	@ID				delete-list
//	@Accept			json
//	@Produce		json
//	@Param			listId	path		integer	true	"list id"
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/lists/{listId} [delete]
func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.DeleteList(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
