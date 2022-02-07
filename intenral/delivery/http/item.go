package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lov3allmy/todo-app/intenral/domain"
	"github.com/lov3allmy/todo-app/intenral/service"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list_id param")
		return
	}

	var input domain.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	itemID, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"item_id": itemID})
}

func (h *Handler) readAllItems(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid list_id param")
		return
	}

	items, err := h.services.TodoItem.ReadAll(userID, listID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllItemsResponse{Data: items})
}

func (h *Handler) readItemByID(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item_id param")
		return
	}

	item, err := h.services.TodoItem.ReadByID(userID, itemID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, item)
}
func (h *Handler) updateItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item_id param")
		return
	}

	var input service.UpdateItemInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userID, itemID, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid item_id param")
		return
	}

	if err := h.services.TodoItem.Delete(userID, itemID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
