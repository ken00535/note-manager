package main

import (
	"note-manager/pkg/app/note/delivery"
	"note-manager/pkg/app/note/repository"
	"note-manager/pkg/app/note/usecase"
	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/db"
	route "note-manager/pkg/infra/http"
	"note-manager/pkg/infra/logger"
)

func main() {
	log := logger.New()
	config.Init(log)
	db.Init(log)
	r := route.NewRoute()
	rg := r.Group("/")
	repo := repository.NewNoteRepository()
	us := usecase.NewNoteUsecase(repo)
	delivery.NewDeliveryHandler(rg, us)
	r.Run(":9300")
}
