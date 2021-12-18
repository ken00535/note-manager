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
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()
	config.Init()
	db.Init(log)
	apiRoute := route.NewRoute()
	apiRoute.LoadHTMLGlob("dist/note-manager/*.html")
	apiRoute.Static("/static", "dist/note-manager")
	apiRoute.Use(gin.Logger())
	apiRoute.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	apiRouteGroup := apiRoute.Group("/api")
	var authDeliver _auth_delivery.Delivery
	{
		repo := _auth_repository.NewAuthRepository()
		us := _auth_usecase.NewAuthUsecase(repo)
		authDeliver = _auth_delivery.NewAuthDelivery(us)
		apiRouteGroup.POST("/login", authDeliver.Login)
		apiRouteGroup.Use(authDeliver.ValidateAuthorization)
	}
	{
		repo := _note_repository.NewNoteRepository()
		us := _note_usecase.NewNoteUsecase(repo)
		deliver := _note_delivery.NewDeliveryHandler(us)
		apiRouteGroup.GET("/notes", deliver.GetNotes)
		apiRouteGroup.GET("/tags", deliver.GetTags)
		apiRouteGroup.Use(authDeliver.ValidatePermission)
		apiRouteGroup.POST("/notes", deliver.AddNote)
		apiRouteGroup.PUT("/notes/:id", deliver.EditNote)
		apiRouteGroup.DELETE("/notes/:id", deliver.DeleteNote)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		apiRoute.Run(":9300")
		wg.Done()
	}()
	go func() {
		apiRoute.RunTLS(":443", "./certs/server.crt", "./certs/server.key")
		wg.Done()
	}()
	wg.Wait()
}
