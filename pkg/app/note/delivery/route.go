package delivery

import (
	"net/http"
	"sync"

	"note-manager/pkg/app/note/usecase"
	"note-manager/pkg/domain/note"
	"note-manager/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

var (
	log  logger.Logger
	once sync.Once
)

// Handler is delivery handler
type Handler struct {
	Usecase usecase.Usecase
}

// NewDeliveryHandler new a delivery
func NewDeliveryHandler(r *gin.RouterGroup, us usecase.Usecase) {
	once.Do(func() {
		log = logger.New()
	})
	handler := Handler{
		Usecase: us,
	}
	r.GET("/api/notes", handler.getNotes)
	r.POST("/api/notes", handler.postNotes)
	r.PUT("/api/notes/:id", handler.putNote)
	r.DELETE("/api/notes/:id", handler.deleteNote)
}

func (h *Handler) getNotes(ctx *gin.Context) {
	type Response struct {
		ID      string `json:"id"`
		Content string `json:"content"`
		Comment string `json:"comment"`
	}
	notes, _ := h.Usecase.GetNotes()
	var resp []Response
	for _, n := range notes {
		r := Response{
			ID:      n.ID,
			Content: n.Content,
			Comment: n.Comment,
		}
		resp = append(resp, r)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) postNotes(ctx *gin.Context) {
	type Request struct {
		Content string `json:"content"`
		Comment string `json:"comment"`
	}
	var reqs []Request
	if err := ctx.ShouldBindJSON(&reqs); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var notes []note.Note
	for _, r := range reqs {
		notes = append(notes, note.Note{
			Content: r.Content,
			Comment: r.Comment,
		})
	}
	h.Usecase.AddNotes(notes)
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handler) putNote(ctx *gin.Context) {
	type Request struct {
		Content string `json:"content"`
		Comment string `json:"comment"`
	}
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	noteID := ctx.Param("id")
	n := note.Note{
		ID:      noteID,
		Content: req.Content,
		Comment: req.Comment,
	}
	h.Usecase.UpdateNote(n)
	ctx.JSON(http.StatusOK, nil)
}

func (h *Handler) deleteNote(ctx *gin.Context) {
	noteID := ctx.Param("id")
	h.Usecase.DeleteNote(noteID)
	ctx.JSON(http.StatusOK, nil)
}
