package handler

import (
	"net/http"
	"strconv"

	"github.com/Cheyzie/golang-test/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetFeedbackById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Error(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	feedback, err := h.service.GetFeedbackById(id)

	if err != nil {
		logrus.Error(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, feedback)
}

func (h *Handler) GetAllFeedbacks(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "NULL")
	offset := ctx.DefaultQuery("offset", "0")

	if limit == "NULL" {
		feedbacks, err := h.service.GetAllFeedbacks()
		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, "somthing went wrong", err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"feedbacks": feedbacks})
		return
	} else {
		feedbacks, err := h.service.GetAllFeedbacksPaginate(limit, offset)

		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, "bad uri", err)
			return
		}

		ctx.JSON(http.StatusOK, feedbacks)
	}
}

func (h *Handler) CreateFeedback(ctx *gin.Context) {
	var input model.Feedback
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusUnprocessableEntity, "Invalid input. Expected customer_name, email, feedback_text and source", err)
		return
	}

	id, err := h.service.CreateFeedback(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid data", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
