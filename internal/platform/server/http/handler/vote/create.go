package vote

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-vote/internal/creating"
	"github.com/rfdez/voting-vote/internal/errors"
	"github.com/rfdez/voting-vote/kit/command"
)

type createRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := commandBus.Dispatch(ctx, creating.NewVoteCommand(
			ctx.Param("poll_id"),
			ctx.Param("option_id"),
			req.UserID,
		)); err != nil {
			switch {
			case errors.IsWrongInput(err):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			default:
				ctx.Status(http.StatusInternalServerError)
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
