package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MuhammadZulkiflii/CRUD-MongoDB/repository"
	"github.com/MuhammadZulkiflii/CRUD-MongoDB/route"

	"go.mongodb.org/mongo-driver/mongo"
)

func shutdown(ctx context.Context, srv *http.Server, db *mongo.Client) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("failed shutdown server")
		return
	}
	if err := db.Disconnect(ctx); err != nil {
		log.Println("failed shutdown database")
		return
	}
	log.Println("shutdown server")
}
func main() {
	db, repo := repository.NewRepository()
	srv := route.NewRoute(repo)
	go func() {
		log.Printf("run server on PORT %s\n", os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	shutdown(ctx, srv, db)
}
