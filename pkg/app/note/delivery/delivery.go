package delivery

import (
	"net/http"
	"strconv"
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
func NewDeliveryHandler(us usecase.Usecase) Handler {
	once.Do(func() {
		log = logger.New()
	})
	handler := Handler{
		Usecase: us,
	}
	return handler
}

// GetNotes get notes
func (h *Handler) GetNotes(ctx *gin.Context) {
	type Response struct {
		ID      string   `json:"id" binding:"alphanum"`
		Content string   `json:"content" binding:"required"`
		Comment string   `json:"comment"`
		Tags    []string `json:"tags"`
	}
	searchKw := ctx.Query("kw")
	tag := ctx.Query("tag")
	pageStr := ctx.Query("page")
	page, _ := strconv.Atoi(pageStr)
	notes, err := h.Usecase.GetNotes(searchKw, tag, page)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := []Response{}
	for _, n := range notes {
		if n.Tags == nil {
			n.Tags = []string{}
		}
		r := Response{
			ID:      n.ID,
			Content: n.Content,
			Comment: n.Comment,
			Tags:    n.Tags,
		}
		resp = append(resp, r)
	}
	ctx.JSON(http.StatusOK, resp)
}

// AddNote add note to db
func (h *Handler) AddNote(ctx *gin.Context) {
	type Request struct {
		Content string   `json:"content" binding:"required"`
		Comment string   `json:"comment"`
		Tags    []string `json:"tags"`
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
			Tags:    r.Tags,
		})
	}
	ids, _ := h.Usecase.AddNotes(notes)
	type Response struct {
		ID string `json:"id"`
	}
	res := Response{ID: ids[0]}
	ctx.JSON(http.StatusOK, res)
}

// EditNote edit note
func (h *Handler) EditNote(ctx *gin.Context) {
	type Request struct {
		Content string   `json:"content" binding:"required"`
		Comment string   `json:"comment"`
		Tags    []string `json:"tags"`
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
		Tags:    req.Tags,
	}
	h.Usecase.UpdateNote(n)
	ctx.JSON(http.StatusOK, nil)
}

// DeleteNote delete note
func (h *Handler) DeleteNote(ctx *gin.Context) {
	noteID := ctx.Param("id")
	h.Usecase.DeleteNote(noteID)
	ctx.JSON(http.StatusOK, nil)
}
