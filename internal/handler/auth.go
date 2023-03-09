package handler

import (
	"net/http"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(ctx *gin.Context) {
	var input model.SigninInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			model.Error{Message: "Invalid input. Expected email and password"},
		)
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid input. Expected email and password", err)
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}
	ctx.JSON(http.StatusOK, token)

}

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid input. Expected username, email and password", err)
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid data", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
