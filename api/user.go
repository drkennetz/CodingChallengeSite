package api

import (
	"net/http"
	"time"

	db "github.com/drkennetz/codingchallenge/db/sqlc"
	"github.com/drkennetz/codingchallenge/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// createUserRequest uses validation to ensure fields are correct
type createUserRequest struct {
	AccountID int64 `json:"account_id" binding:"required"`
	AdminUser bool `json:"admin_user" binding:"required"`
	Username  string `json:"username" binding:"required,alphanum"`
	Password  string `json:"password" binding:"required,min=9"`
	Grade db.DevLevel `json:"grade" binding:"required"`
}

type createUserResponse struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	AdminUser bool      `json:"admin_user"`
	Username  string    `json:"username"`
	Grade     db.DevLevel  `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		AccountID: req.AccountID,
		AdminUser: req.AdminUser,
		Username: req.Username,
		Password: hashedPassword,
		Grade: req.Grade,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse {
		ID: user.ID,
		AccountID: user.AccountID,
		AdminUser: user.AdminUser,
		Username: user.Username,
		Grade: user.Grade,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}