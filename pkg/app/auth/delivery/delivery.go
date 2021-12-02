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
	usecase usecase.Usecase
}

// NewAuthDelivery new a delivery
func NewAuthDelivery(us usecase.Usecase) Delivery {
	once.Do(func() {
		log = logger.New()
	})
	handler := Delivery{
		usecase: us,
	}
	return handler
}

// ValidateAuthorization validate authorization
func (h *Delivery) ValidateAuthorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]
	username, err := h.usecase.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Set("username", username)
	ctx.Next()
}

// ValidatePermission validate permission
func (h *Delivery) ValidatePermission(ctx *gin.Context) {
	username := ctx.GetString("username")
	err := h.usecase.ValidatePermission(username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}

// Login login to service
func (h *Delivery) Login(ctx *gin.Context) {
	type Request struct {
		Username string `json:"username" binding:"alphanum"`
		Password string `json:"password" binding:"alphanum"`
	}
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
