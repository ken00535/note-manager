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
	r.GET("/api/notes", handler.getNotes)
	r.GET("/api/notes/labels/:label", handler.getNotes)
}

func (h *Handler) getNotes(ctx *gin.Context) {
	type Response struct {
		Content string `json:"content"`
		Comment string `json:"comment"`
	}
	notes, _ := h.Usecase.GetNotes()
	var resp []Response
	for _, n := range notes {
		r := Response{
			Content: n.Content,
			Comment: n.Comment,
		}
		resp = append(resp, r)
	}
	ctx.JSON(http.StatusOK, resp)
}
