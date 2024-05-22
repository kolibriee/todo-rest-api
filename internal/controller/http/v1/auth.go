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
	id, err := h.services.Autorization.SignUp(input)
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
	accessToken, refreshToken, err := h.services.Autorization.SignIn(c.ClientIP(), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", refreshToken, 0, "/", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	clientIP := c.ClientIP()
	accessToken, newRefreshToken, err := h.services.Autorization.Refresh(refreshToken, clientIP)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("refreshToken", newRefreshToken, 0, "/", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
	})
}
