package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/lov3allmy/todo-app/intenral/service"
	"net/http"
	"strconv"
)

func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	var input domain.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) readAllLists(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.ReadAll(userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

func (h *Handler) readListByID(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.ReadByID(userID, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input service.UpdateListInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userID, id, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.TodoList.Delete(userID, id); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
