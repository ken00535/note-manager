package main

import (
	"note-manager/pkg/infra/db"
	route "note-manager/pkg/infra/http"
	"note-manager/pkg/infra/logger"
)

func main() {
	log := logger.New()
	db.Connect(log)
	r := route.NewRoute()
	r.Run(":9300")
}
