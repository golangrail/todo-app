package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type getAllListsResponse struct {
	Data []domain.TodoList `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
