package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
)

// .
//
//	@Summary		SignUp
//	@Tags			auth
//	@Description	create account
//	@ID				create-account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.User	true	"account info"
//	@Success		200		{integer}	integer		1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input domain.User

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

// .
//
//	@Summary		SignIn
//	@Tags			auth
//	@Description	authenticate user
//	@ID				authenticate-user
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.SignInUserInput	true	"login password"
//	@Success		200		{string}	string					"token"
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignInUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.services.Autorization.GenerateToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
