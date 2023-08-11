package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tryrest "github.com/kolibri7557/try-rest-api"
)

func (h *Handler) signUp(c *gin.Context) {
	var input tryrest.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
