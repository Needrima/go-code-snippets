package main

import (
	databaseadapter "go-code-snippets/internal/adapters/database-adapter"
	"go-code-snippets/internal/adapters/httpadapter"
	"go-code-snippets/internal/adapters/routes"
	"go-code-snippets/internal/core/helper"
	"go-code-snippets/internal/core/processor"
	"log"
	"net/http"
	"os"
)

func main() {
	helper.LoadEnv("./app.env")
	dbCollection := databaseadapter.ConnectToDB()

	databaseadapter := databaseadapter.NewDBAdapter(dbCollection)
	processor := processor.NewProcessor(databaseadapter)
	httpadapter := httpadapter.NewHandler(processor)

	router := routes.SetupRouter(httpadapter)

	port := os.Getenv("PORT")
	log.Printf("listening on port: %s............\n", port)
	http.ListenAndServe(":"+port, router)
}
