package api

import (
	db "backend/db/sqlc"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) listTutorials(ctx *gin.Context) {
	tutorials, err := server.store.GetTutorials(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tutorials)
}

type getTutorialIdRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) listUserTutorials(ctx *gin.Context) {
	var requestData getTutorialIdRequest

	if err := ctx.ShouldBindUri(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	tutorials, err := server.store.GetTutorialsByUser(ctx, requestData.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tutorials)
}

type createTutorialRequest struct {
	Title    string `json:"title" binding:"required"`
	Material string `json:"material" binding:"required"`
	UserId   int32  `json:"user_Id" binding:"required"`
}

func (server *Server) createTutorial(ctx *gin.Context) {
	var requestData createTutorialRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	arg := db.CreateTutorialParams{
		Title:    requestData.Title,
		Material: requestData.Material,
		UserID:   requestData.UserId,
	}

	tutorial, err := server.store.CreateTutorial(ctx, arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tutorial)
}

type updateTutorialRequest struct {
	Id       int32  `json:"id" binding:"required"`
	Material string `json:"material" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

func (server *Server) updateTutorial(ctx *gin.Context) {
	var requestData updateTutorialRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	var requestId getTutorialIdRequest

	if err := ctx.ShouldBindUri(&requestId); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	arg := db.UpdateTutorialParams{
		ID:       requestId.ID,
		Material: requestData.Material,
		Title:    requestData.Title,
	}

	tutorial, err := server.store.UpdateTutorial(ctx, arg)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tutorial)
}

type getTutorialRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTutorial(ctx *gin.Context) {
	var requestData getTutorialRequest

	if err := ctx.ShouldBindUri(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	tutorial, err := server.store.GetTutorial(ctx, requestData.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tutorial)
}

func (server *Server) deleteTutorial(ctx *gin.Context) {
	var requestData getTutorialRequest

	if err := ctx.ShouldBindUri(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	err := server.store.DeleteTutorial(ctx, requestData.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
