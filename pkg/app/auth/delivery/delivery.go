package delivery

import (
	"net/http"
	"strings"
	"sync"

	"note-manager/pkg/app/auth/usecase"
	"note-manager/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

var (
	log  logger.Logger
	once sync.Once
)

// Delivery is delivery handler
type Delivery struct {
	Router  gin.IRouter
	usecase usecase.Usecase
}

// NewAuthDelivery new a delivery
func NewAuthDelivery(r gin.IRouter, us usecase.Usecase) Delivery {
	once.Do(func() {
		log = logger.New()
	})
	handler := Delivery{
		Router:  r,
		usecase: us,
	}
	r.POST("/api/login", handler.postLogin)
	return handler
}

// ValidateAuthorization validate authorization
func (h *Delivery) ValidateAuthorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]
	err := h.usecase.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (h *Delivery) postLogin(ctx *gin.Context) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// check account and password is correct
	err := h.usecase.ValidateUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	token, err := h.usecase.GetToken(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
