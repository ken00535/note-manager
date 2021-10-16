package delivery

import (
	"net/http"

	"note-manager/pkg/app/note/usecase"

	"github.com/gin-gonic/gin"
)

// Handler is delivery handler
type Handler struct {
	Usecase usecase.Usecase
}

// NewDeliveryHandler new a delivery
func NewDeliveryHandler(r *gin.RouterGroup, us usecase.Usecase) {
	handler := Handler{
		Usecase: us,
	}
	r.GET("/notes/labels/:label", handler.getSome)
}

func (h *Handler) getSome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
