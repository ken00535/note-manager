package main

import (
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
	r := route.NewRoute()
	var validateAuthorization func(ctx *gin.Context)
	rg := r.Group("/")
	{
		repo := _auth_repository.NewAuthRepository()
		us := _auth_usecase.NewAuthUsecase(repo)
		deliver := _auth_delivery.NewAuthDelivery(rg, us)
		validateAuthorization = deliver.ValidateAuthorization
	}
	rg.Use(validateAuthorization)
	{
		repo := _note_repository.NewNoteRepository()
		us := _note_usecase.NewNoteUsecase(repo)
		_note_delivery.NewDeliveryHandler(rg, us)
	}
	r.Run(":9300")
}
