package main

import (
	databaseadapter "go-code-snippets/internal/adapters/database-adapter"
	httpadapter "go-code-snippets/internal/adapters/http-adapter"
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
	databaseAdapter := databaseadapter.NewDBAdapter(dbCollection)
	processor := processor.NewProcessor(databaseAdapter)
	httpAdapter := httpadapter.NewHandler(processor)

	router := routes.SetupRouter(httpAdapter)

	port := os.Getenv("PORT")
	log.Printf("listening on port: %s............\n", port)
	http.ListenAndServe(":"+port, router)
}
