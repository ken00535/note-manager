package main

import (
	"net/http"
	_auth_delivery "note-manager/pkg/app/auth/delivery"
	_auth_repository "note-manager/pkg/app/auth/repository"
	_auth_usecase "note-manager/pkg/app/auth/usecase"
	_note_delivery "note-manager/pkg/app/note/delivery"
	_note_repository "note-manager/pkg/app/note/repository"
	_note_usecase "note-manager/pkg/app/note/usecase"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/db"
	route "note-manager/pkg/infra/http"
	"note-manager/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()
	config.Init(log)
	db.Init(log)
	apiRoute := route.NewRoute()
	apiRoute.LoadHTMLGlob("dist/note-manager/*.html")
	apiRoute.Static("/static", "dist/note-manager")
	apiRoute.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	apiRouteGroup := apiRoute.Group("/api")
	{
		repo := _auth_repository.NewAuthRepository()
		us := _auth_usecase.NewAuthUsecase(repo)
		deliver := _auth_delivery.NewAuthDelivery(apiRouteGroup, us)
		apiRouteGroup.Use(deliver.ValidateAuthorization)
	}
	{
		repo := _note_repository.NewNoteRepository()
		us := _note_usecase.NewNoteUsecase(repo)
		_note_delivery.NewDeliveryHandler(apiRouteGroup, us)
	}
	apiRoute.Run(":9300")
}
