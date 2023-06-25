package main

import (
	"assisment/internal/boot"
	"assisment/internal/httpservice"
	"assisment/pkg/buffer"
	"assisment/pkg/database"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	_, cancel := context.WithCancel(boot.NewContext(context.Background()))
	defer cancel()
	fmt.Println("started app")
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateLogTable()
	if err != nil {
		log.Fatal(err)
	}
	logBuffer := buffer.NewLogBuffer()
	router := mux.NewRouter()
	httpservice.RouteHanler(router, logBuffer)

	go logBuffer.StartProcessing(db)
	http.ListenAndServe(":8000", router)
}
